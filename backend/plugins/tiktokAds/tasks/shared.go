/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/tiktokAds/models"
	"io"
	"net/http"
	"strconv"
)

//func CreateRawDataSubTaskArgs(taskCtx plugin.SubTaskContext, rawTable string) (*api.RawDataSubTaskArgs, *TiktokAdsTaskData) {
//	data := taskCtx.GetData().(*TiktokAdsTaskData)
//	filteredData := *data
//	filteredData.Options = &TiktokAdsOptions{}
//	*filteredData.Options = *data.Options
//	var params = TiktokAdsApiParams{
//		ConnectionId: data.Options.ConnectionId,
//		ProjectKey:   data.Options.ProjectKey,
//		HotspotKey:   data.Options.HotspotKey,
//	}
//	rawDataSubTaskArgs := &api.RawDataSubTaskArgs{
//		Ctx:    taskCtx,
//		Params: params,
//		Table:  rawTable,
//	}
//	return rawDataSubTaskArgs, &filteredData
//}

func GetRawMessageArrayFromResponse(res *http.Response) ([]json.RawMessage, errors.Error) {
	var data struct {
		Data struct {
			PageInfo models.PageInfo   `json:"page_info"`
			List     []json.RawMessage `json:"list"`
		} `json:"data"`
	}
	err := api.UnmarshalResponse(res, &data)
	return data.Data.List, err
}

func GetTotalPages(res *http.Response, args *api.ApiCollectorArgs) (int, errors.Error) {
	var data struct {
		Data struct {
			PageInfo models.PageInfo `json:"page_info"`
		} `json:"data"`
	}
	err := api.UnmarshalResponse(res, &data)
	if err != nil {
		return 0, err
	}
	return data.Data.PageInfo.TotalPage, nil
}

func prepareUpdate(recordType string, recordIds []uint64, data *TiktokAdsTaskData, operate string) {
	count := 0
	recordsQuery := make([]string, 0)
	length := len(recordIds)
	idsFieldName := fmt.Sprintf("%s_ids", recordType)
	for _, recordId := range recordIds {
		count++
		recordsQuery = append(recordsQuery, fmt.Sprintf("%d", recordId))
		if count%20 == 0 || count == length {
			payload := map[string]interface{}{
				"advertiser_id":    data.Options.AdvertiserID,
				idsFieldName:       recordsQuery,
				"operation_status": operate,
			}
			updateStatus(recordType+"/status", payload, data.ApiClient)
			recordsQuery = make([]string, 0)
		}
	}
}

func modifyField(modifyBudgetMap map[models.TiktokAdsRule][]models.TiktokAdsAdGroupReport, data *TiktokAdsTaskData) {
	for rule, elems := range modifyBudgetMap {
		for _, elem := range elems {
			valueToRevise := 0.0
			rate := 1 + rule.ValueToRevise
			payload := map[string]interface{}{
				"advertiser_id": data.Options.AdvertiserID,
				"adgroup_id":    strconv.FormatUint(elem.AdgroupId, 10),
			}
			switch rule.FieldToRevise {
			case "budget":
				valueToRevise = elem.Budget * rate
			case "conversion_bid_price":
				if elem.Bid == "-" {
					break
				}
				bid, err := strconv.ParseFloat(elem.Bid, 64)
				if err != nil {
					fmt.Println(err)
					return
				}
				valueToRevise = bid * rate
			}
			payload[rule.FieldToRevise] = valueToRevise
			updateStatus("adgroup", payload, data.ApiClient)
		}
	}
}

func updateStatus(recordType string, payload map[string]interface{}, apiClient *api.ApiAsyncClient) {
	url := fmt.Sprintf("https://ads.tiktok.com/open_api/v1.3/%s/update/", recordType)
	res, err := apiClient.Post(url, nil, payload, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err1 := io.ReadAll(res.Body)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(string(body))
}
