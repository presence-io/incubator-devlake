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
	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/tiktokAds/models"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strconv"
	"time"
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

// prepareUpdate 准备更新记录
// recordType: 记录类型
// recordIds: 记录ID列表
// data: Tiktok广告任务数据
// operate: 操作
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

// modifyField 修改字段
// modifyBudgetMap: 预算修改映射
// data: Tiktok广告任务数据
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

			payload[rule.FieldToRevise] = fmt.Sprintf("%.2f", valueToRevise)
			updateStatus("adgroup", payload, data.ApiClient)
		}
	}
}

// updateStatus 更新状态
// recordType: 记录类型
// payload: 请求负载数据
// apiClient: API客户端
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
	jsonBody := make(map[string]interface{})
	err1 = json.Unmarshal(body, &jsonBody)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(err1)
	if jsonBody["message"] == "OK" {
		jsonBody["中文提示"] = "操作成功"
	} else {
		jsonBody["中文提示"] = "操作失败"
	}

	feishuPayload := map[string]interface{}{
		"req": payload,
		"res": jsonBody,
	}
	jsonPayload, err1 := json.Marshal(&feishuPayload)
	if err1 != nil {
		fmt.Println(err1)
	}
	feishuNotify(string(jsonPayload), err, apiClient)
}

// feishuNotify 飞书通知
// payload: 请求负载数据
// resMsg: 响应消息
// err: 错误
// apiClient: API客户端
func feishuNotify(resMsg string, err error, apiClient *api.ApiAsyncClient) {
	url := "https://open.feishu.cn/open-apis/bot/v2/hook/8676a3dd-f9f0-49d8-9ac4-6edfbc1b7fae"
	req := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]string{
			"text": fmt.Sprintf("%s, 错误：%v", resMsg, err),
		},
	}
	res, err := apiClient.Post(url, nil, req, nil)
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

// calculateRule 根据规则条件计算是否满足条件
// conditions: 规则条件列表
// reportValueMap: 报告值映射
// taskCtx: 任务上下文
func calculateRule(conditions []*models.TiktokAdsRuleCondition, reportValueMap map[string]interface{}, taskCtx plugin.SubTaskContext) bool {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	db := taskCtx.GetDal()
	var float64Value float64
	referValue := 0.0

	for _, condition := range conditions {
		if condition.TimeRange != 0 {
			today := reportValueMap["stat_time_day"].(string)
			todayTime, err := time.Parse("2006-01-02 00:00:00", today)
			if err != nil {
				panic(err)
			}
			if condition.TimeRange%24 == 0 {
				days := condition.TimeRange / 24
				compareDate := todayTime.AddDate(0, 0, int(-days)).Format("2006-01-02 00:00:00")
				clauses := []dal.Clause{
					dal.From(&models.TiktokAdsAdGroupReport{}),
					dal.Where("adgroup_id = ? and stat_time_day = ? and connection_id = ?", reportValueMap["id"], compareDate, data.Options.ConnectionId),
				}
				compareAdGroupReport := &models.TiktokAdsAdGroupReport{}
				err = db.First(compareAdGroupReport, clauses...)
				if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					panic(err)
				}
				if compareAdGroupReport == nil {
					return false
				}
				compareReportValueMap := make(map[string]interface{})
				err = errors.Convert(mapstructure.Decode(&compareAdGroupReport.TiktokAdsReportCommon, &compareReportValueMap))
				if err != nil {
					panic(err)
				}
				reportValueMap = compareReportValueMap
			}
		}

		// 如果报告中不包含条件中的字段，则直接返回 false
		value, ok := reportValueMap[condition.FieldName]
		if !ok {
			panic(fmt.Sprintf("report does not contain %s", condition.FieldName))
		}

		float64Value = convertToFloat64(value)

		if condition.FieldRelationName != "" && condition.FieldRelationRate != 0 {
			value, ok := reportValueMap[condition.FieldRelationName]
			if !ok {
				panic(fmt.Sprintf("in condition, %s has wrong value %f", condition.FieldRelationName, condition.FieldRelationRate))
			}
			referValue = convertToFloat64(value) * condition.FieldRelationRate
		} else {
			referValue = condition.FieldValue
		}

		if !compareValues(float64Value, referValue, condition.Operator) {
			return false
		}
	}
	return true
}

// convertToFloat64 将值转换为 float64 类型，如果无法转换则返回 0.0
func convertToFloat64(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case int:
		return float64(v)
	default:
		return 0.0
	}
}

// compareValues 根据操作符比较两个值
func compareValues(value1, value2 float64, operator string) bool {
	switch operator {
	case ">":
		return value1 > value2
	case "<":
		return value1 < value2
	case "=":
		return value1 == value2
	default:
		return false
	}
}
