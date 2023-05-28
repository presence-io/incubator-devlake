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

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

const RAW_REPORT_TABLE = "tiktokads_report"

var _ plugin.SubTaskEntryPoint = CollectReport

func CollectReport(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	collector, err := api.NewApiCollector(api.ApiCollectorArgs{
		RawDataSubTaskArgs: api.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TiktokAdsApiParams{
				ConnectionId: data.Options.ConnectionId,
				AdvertiserId: data.Options.AdvertiserID,
			},
			Table: RAW_REPORT_TABLE,
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
			query.Set("report_type", data.Options.ReportType)
			query.Set("data_level", data.Options.DataLevel)
			query.Set("start_date", data.Options.StartDate)
			query.Set("end_date", data.Options.EndDate)
			for _, metric := range data.Options.Metrics {
				query.Add("metrics", metric)
			}
			for _, dimension := range data.Options.Dimensions {
				query.Add("dimensions", dimension)
			}
			query.Set("query_lifetime", strconv.FormatBool(data.Options.QueryLifetime))
			query.Set("order_field", data.Options.OrderField)
			// OrderType
			query.Set("order_type", data.Options.OrderType)
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

var CollectReportMeta = plugin.SubTaskMeta{
	Name:             "CollectReport",
	EntryPoint:       CollectReport,
	EnabledByDefault: true,
	Description:      "Collect Report data from Tiktokads api",
	DomainTypes:      []string{},
}

var campaignMetrics = []string{
	"spend",
	"cpc",
	"cpm",
	"impressions",
	"gross_impressions",
	"clicks",
	"ctr",
	"reach",
	"cost_per_1000_reached",
	"conversion",
	"cost_per_conversion",
	"conversion_rate",
	"real_time_conversion",
	"real_time_cost_per_conversion",
	"real_time_conversion_rate",
	"result",
	"cost_per_result",
	"result_rate",
	"real_time_result",
	"real_time_cost_per_result",
	"real_time_result_rate",
	"secondary_goal_result",
	"cost_per_secondary_goal_result",
	"secondary_goal_result_rate",
	"frequency",
	"currency",
	"campaign_name",
	"objective_type",
	"split_test",
	"campaign_budget",
	"campaign_dedicate_type",
	"app_promotion_type",
	"onsite_download_start",
	"cost_per_onsite_download_start",
	"onsite_download_start_rate",
}
var adGroupMetrics = []string{
	"spend",
	"cpc",
	"cpm",
	"impressions",
	"gross_impressions",
	"clicks",
	"ctr",
	"reach",
	"cost_per_1000_reached",
	"conversion",
	"cost_per_conversion",
	"conversion_rate",
	"real_time_conversion",
	"real_time_cost_per_conversion",
	"real_time_conversion_rate",
	"result",
	"cost_per_result",
	"result_rate",
	"real_time_result",
	"real_time_cost_per_result",
	"real_time_result_rate",
	"secondary_goal_result",
	"cost_per_secondary_goal_result",
	"secondary_goal_result_rate",
	"frequency",
	"currency",
	"campaign_name",
	"campaign_id",
	"objective_type",
	"split_test",
	"campaign_budget",
	"campaign_dedicate_type",
	"app_promotion_type",
	"adgroup_name",
	"placement_type",
	"promotion_type",
	"opt_status",
	"dpa_target_audience_type",
	"budget",
	"smart_target",
	"pricing_category",
	"bid_strategy",
	"bid",
	"aeo_type",
	"tt_app_id",
	"tt_app_name",
	"mobile_app_id",
	"onsite_download_start",
	"cost_per_onsite_download_start",
	"onsite_download_start_rate",
}

var adMetrics = []string{
	"spend",
	"cpc",
	"cpm",
	"impressions",
	"gross_impressions",
	"clicks",
	"ctr",
	"reach",
	"cost_per_1000_reached",
	"conversion",
	"cost_per_conversion",
	"conversion_rate",
	"real_time_conversion",
	"real_time_cost_per_conversion",
	"real_time_conversion_rate",
	"result",
	"cost_per_result",
	"result_rate",
	"real_time_result",
	"real_time_cost_per_result",
	"real_time_result_rate",
	"secondary_goal_result",
	"cost_per_secondary_goal_result",
	"secondary_goal_result_rate",
	"frequency",
	"currency",
	"campaign_name",
	"campaign_id",
	"objective_type",
	"split_test",
	"campaign_budget",
	"campaign_dedicate_type",
	"app_promotion_type",
	"adgroup_name",
	"adgroup_id",
	"placement_type",
	"promotion_type",
	"opt_status",
	"dpa_target_audience_type",
	"budget",
	"smart_target",
	"pricing_category",
	"bid_strategy",
	"bid",
	"aeo_type",
	"ad_name",
	"ad_text",
	"call_to_action",
	"tt_app_id",
	"tt_app_name",
	"mobile_app_id",
	"image_mode",
	"onsite_download_start",
	"cost_per_onsite_download_start",
	"onsite_download_start_rate",
}
