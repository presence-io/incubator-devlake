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

package main

import (
	"github.com/apache/incubator-devlake/core/runner"
	"github.com/apache/incubator-devlake/plugins/tiktokAds/impl"
	"github.com/spf13/cobra"
)

// Export a variable named PluginEntry for Framework to search and load
var PluginEntry impl.TiktokAds //nolint

// standalone mode for debugging
func main() {
	//cmd := &cobra.Command{Use: "github"}
	//connectionId := cmd.Flags().Uint64P("connectionId", "c", 0, "github connection id")
	//owner := cmd.Flags().StringP("owner", "o", "", "github owner")
	//repo := cmd.Flags().StringP("repo", "r", "", "github repo")
	//timeAfter := cmd.Flags().StringP("timeAfter", "a", "", "collect data that are created after specified time, ie 2006-05-06T07:08:09Z")
	//_ = cmd.MarkFlagRequired("connectionId")
	//_ = cmd.MarkFlagRequired("owner")
	//_ = cmd.MarkFlagRequired("repo")
	//
	//prType := cmd.Flags().String("prType", "type/(.*)$", "pr type")
	//prComponent := cmd.Flags().String("prComponent", "component/(.*)$", "pr component")
	//prBodyClosePattern := cmd.Flags().String("prBodyClosePattern", "(?mi)(fix|close|resolve|fixes|closes|resolves|fixed|closed|resolved)[\\s]*.*(((and )?(#|https:\\/\\/github.com\\/%s\\/issues\\/)\\d+[ ]*)+)", "pr body close pattern")
	//issueSeverity := cmd.Flags().String("issueSeverity", "severity/(.*)$", "issue severity")
	//issuePriority := cmd.Flags().String("issuePriority", "^(highest|high|medium|low)$", "issue priority")
	//issueComponent := cmd.Flags().String("issueComponent", "component/(.*)$", "issue component")
	//issueTypeBug := cmd.Flags().String("issueTypeBug", "^(bug|failure|error)$", "issue type bug")
	//issueTypeIncident := cmd.Flags().String("issueTypeIncident", "", "issue type incident")
	//issueTypeRequirement := cmd.Flags().String("issueTypeRequirement", "^(feat|feature|proposal|requirement)$", "issue type requirement")
	//deploymentPattern := cmd.Flags().StringP("deployment", "", "", "deployment pattern")
	//productionPattern := cmd.Flags().StringP("production", "", "", "production pattern")
	//
	//cmd.Run = func(cmd *cobra.Command, args []string) {
	//	runner.DirectRun(cmd, args, PluginEntry, map[string]interface{}{
	//		"connectionId": *connectionId,
	//		"owner":        *owner,
	//		"repo":         *repo,
	//		"timeAfter":    *timeAfter,
	//		"transformationRules": map[string]interface{}{
	//			"prType":               *prType,
	//			"prComponent":          *prComponent,
	//			"prBodyClosePattern":   *prBodyClosePattern,
	//			"issueSeverity":        *issueSeverity,
	//			"issuePriority":        *issuePriority,
	//			"issueComponent":       *issueComponent,
	//			"issueTypeBug":         *issueTypeBug,
	//			"issueTypeIncident":    *issueTypeIncident,
	//			"issueTypeRequirement": *issueTypeRequirement,
	//			"deploymentPattern":    *deploymentPattern,
	//			"productionPattern":    *productionPattern,
	//		},
	//	})
	//}
	cmd := &cobra.Command{Use: "tiktokAds"}

	// TODO add your cmd flag if necessary
	//connectionId := cmd.Flags().Uint64P("connectionId", "c", 0, "github connection id")
	//advertiser_id := cmd.Flags().StringP("owner", "o", "", "github owner")
	//dimensions := cmd.Flags().StringP("repo", "r", "", "github repo")
	campaignMetrics := []string{
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
	//adGroupMetrics := []string{
	//	"spend",
	//	"cpc",
	//	"cpm",
	//	"impressions",
	//	"gross_impressions",
	//	"clicks",
	//	"ctr",
	//	"reach",
	//	"cost_per_1000_reached",
	//	"conversion",
	//	"cost_per_conversion",
	//	"conversion_rate",
	//	"real_time_conversion",
	//	"real_time_cost_per_conversion",
	//	"real_time_conversion_rate",
	//	"result",
	//	"cost_per_result",
	//	"result_rate",
	//	"real_time_result",
	//	"real_time_cost_per_result",
	//	"real_time_result_rate",
	//	"secondary_goal_result",
	//	"cost_per_secondary_goal_result",
	//	"secondary_goal_result_rate",
	//	"frequency",
	//	"currency",
	//	"campaign_name",
	//	"campaign_id",
	//	"objective_type",
	//	"split_test",
	//	"campaign_budget",
	//	"campaign_dedicate_type",
	//	"app_promotion_type",
	//	"adgroup_name",
	//	"placement_type",
	//	"promotion_type",
	//	"opt_status",
	//	"dpa_target_audience_type",
	//	"budget",
	//	"smart_target",
	//	"pricing_category",
	//	"bid_strategy",
	//	"bid",
	//	"aeo_type",
	//	"tt_app_id",
	//	"tt_app_name",
	//	"mobile_app_id",
	//	"onsite_download_start",
	//	"cost_per_onsite_download_start",
	//	"onsite_download_start_rate",
	//}
	//
	//adMetricst := []string{
	//	"spend",
	//	"cpc",
	//	"cpm",
	//	"impressions",
	//	"gross_impressions",
	//	"clicks",
	//	"ctr",
	//	"reach",
	//	"cost_per_1000_reached",
	//	"conversion",
	//	"cost_per_conversion",
	//	"conversion_rate",
	//	"real_time_conversion",
	//	"real_time_cost_per_conversion",
	//	"real_time_conversion_rate",
	//	"result",
	//	"cost_per_result",
	//	"result_rate",
	//	"real_time_result",
	//	"real_time_cost_per_result",
	//	"real_time_result_rate",
	//	"secondary_goal_result",
	//	"cost_per_secondary_goal_result",
	//	"secondary_goal_result_rate",
	//	"frequency",
	//	"currency",
	//	"campaign_name",
	//	"campaign_id",
	//	"objective_type",
	//	"split_test",
	//	"campaign_budget",
	//	"campaign_dedicate_type",
	//	"app_promotion_type",
	//	"adgroup_name",
	//	"adgroup_id",
	//	"placement_type",
	//	"promotion_type",
	//	"opt_status",
	//	"dpa_target_audience_type",
	//	"budget",
	//	"smart_target",
	//	"pricing_category",
	//	"bid_strategy",
	//	"bid",
	//	"aeo_type",
	//	"ad_name",
	//	"ad_text",
	//	"call_to_action",
	//	"tt_app_id",
	//	"tt_app_name",
	//	"mobile_app_id",
	//	"image_mode",
	//	"onsite_download_start",
	//	"cost_per_onsite_download_start",
	//	"onsite_download_start_rate",
	//}
	cmd.Run = func(cmd *cobra.Command, args []string) {
		runner.DirectRun(cmd, args, PluginEntry, map[string]interface{}{
			"connectionId": 1,
			"advertiserId": "7232182199037034497",
			"reportType":   "BASIC",
			"dimensions":   []string{"1739039623689250", "1738056643731457"},
			"metrics":      campaignMetrics,
		})
	}
	runner.RunCmd(cmd)
}
