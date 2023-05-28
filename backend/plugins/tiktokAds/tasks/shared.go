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
	"strings"
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
	payloadRecordsString := `["`
	for _, recordId := range recordIds {
		payloadRecordsString += fmt.Sprintf(`%d",`, recordId)
		count++
		if count%20 == 0 {
			payloadRecordsString = payloadRecordsString + `"]`
			payload := strings.NewReader(fmt.Sprintf(`{
					"advertiser_id": %s,
					"adgroup_ids": %s,
					"operation_status": %s
				}`, data.Options.AdvertiserID, payloadRecordsString, operate))
			updateStatus(recordType+"/status", payload, data.ApiClient)
			payloadRecordsString = `["`
		}
	}
}

func modifyBudget(modifyBudgetMap map[models.TiktokAdsRule]models.TiktokAdsAdGroupReport, data *TiktokAdsTaskData) {
	for rule, budget := range modifyBudgetMap {
		payload := strings.NewReader(fmt.Sprintf(`{
    "advertiser_id": %s,
    "budget": [
        {
            "adgroup_id": %d,
            "budget": %f
        }
    ]
}`, data.Options.AdvertiserID, budget.AdgroupId, budget.Budget+rule.BudgetToRevise))

		updateStatus("adgroup/budget", payload, data.ApiClient)
	}
}

func updateStatus(recordType string, payload *strings.Reader, apiClient *api.ApiAsyncClient) {
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
