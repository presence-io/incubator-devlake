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
	"net/url"
	"strconv"
	"time"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

var _ plugin.SubTaskEntryPoint = CollectAdGroupReport

const RAW_AD_GROUP_REPORT_TABLE = "tiktokads_ad_group_report"

func CollectAdGroupReport(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	now := time.Now() // 获取当前时间
	collector, err := api.NewApiCollector(api.ApiCollectorArgs{
		RawDataSubTaskArgs: api.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TiktokAdsApiParams{
				ConnectionId: data.Options.ConnectionId,
				AdvertiserId: data.Options.AdvertiserID,
			},
			Table: RAW_AD_GROUP_REPORT_TABLE,
		},
		ApiClient:     data.ApiClient,
		Incremental:   false,
		UrlTemplate:   "v1.3/report/integrated/get/",
		GetTotalPages: GetTotalPages,
		PageSize:      50,
		Query: func(reqData *api.RequestData) (url.Values, errors.Error) {

			query := url.Values{}
			query.Set("advertiser_id", data.Options.AdvertiserID)
			query.Set("service_type", data.Options.ServiceType)
			query.Set("report_type", "BASIC")
			query.Set("data_level", "AUCTION_ADGROUP")
			query.Set("start_date", now.Format("2006-01-02"))
			query.Set("end_date", now.Format("2006-01-02"))
			for _, metric := range adGroupMetrics {
				query.Add("metrics", metric)
			}
			//data.Options.Dimensions = []string{"adgroup_id"}
			for _, dimension := range []string{"adgroup_id", "stat_time_day"} {
				query.Add("dimensions", dimension)
			}
			//query.Set("query_lifetime", strconv.FormatBool(data.Options.QueryLifetime))
			query.Set("order_field", "spend")
			// OrderType
			query.Set("order_type", "DESC")
			// Filtering
			// Set the filtering field as a query parameter
			for _, filter := range data.Options.Filtering {
				filterJSON, _ := json.Marshal(filter)
				query.Add("filtering", string(filterJSON))
			}
			//QueryMode
			query.Set("query_mode", data.Options.QueryMode)
			query.Set("page", strconv.Itoa(reqData.Pager.Page))
			query.Set("page_size", strconv.Itoa(reqData.Pager.Size))
			return query, nil
		},
		ResponseParser: GetRawMessageArrayFromResponse,
	})

	if err != nil {
		return err
	}
	return collector.Execute()
}

var CollectAdGroupReportMeta = plugin.SubTaskMeta{
	Name:             "CollectAdGroupReport",
	EntryPoint:       CollectAdGroupReport,
	EnabledByDefault: true,
	Description:      "Collect AdGroupReport data from Tiktokads api",
	DomainTypes:      []string{},
}
