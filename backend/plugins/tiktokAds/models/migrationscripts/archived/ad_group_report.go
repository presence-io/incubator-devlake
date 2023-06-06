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

package archived

import "github.com/apache/incubator-devlake/core/models/migrationscripts/archived"

type TiktokAdsAdGroupReport struct {
	ConnectionId               uint64  `json:"connection_id" gorm:"column:connection_id;autoIncrement:false;primaryKey"`
	StatTimeDay                string  `json:"stat_time_day" gorm:"column:stat_time_day;primaryKey"`
	AdgroupId                  uint64  `json:"adgroup_id,string" gorm:"column:adgroup_id;primaryKey;autoIncrement:false"`
	AdvertiserID               string  `json:"advertiser_id,omitempty" gorm:"column:advertiser_id;primaryKey"`
	AdgroupName                string  `json:"adgroup_name" gorm:"column:adgroup_name"`
	AeoType                    string  `json:"aeo_type" gorm:"column:aeo_type"`
	AppPromotionType           string  `json:"app_promotion_type" gorm:"column:app_promotion_type"`
	Bid                        string  `json:"bid" gorm:"column:bid"`
	BidStrategy                string  `json:"bid_strategy" gorm:"column:bid_strategy"`
	Budget                     float64 `json:"budget,string" gorm:"column:budget"`
	CampaignBudget             float64 `json:"campaign_budget,string" gorm:"column:campaign_budget"`
	CampaignDedicateType       string  `json:"campaign_dedicate_type" gorm:"column:campaign_dedicate_type"`
	CampaignId                 uint64  `json:"campaign_id,string" gorm:"column:campaign_id"`
	CampaignName               string  `json:"campaign_name" gorm:"column:campaign_name"`
	Clicks                     int     `json:"clicks,string" gorm:"column:clicks"`
	Conversion                 int     `json:"conversion,string" gorm:"column:conversion"`
	ConversionRate             float64 `json:"conversion_rate,string" gorm:"column:conversion_rate"`
	CostPer1000Reached         float64 `json:"cost_per_1000_reached,string" gorm:"column:cost_per_1000_reached"`
	CostPerConversion          float64 `json:"cost_per_conversion,string" gorm:"column:cost_per_conversion"`
	CostPerOnsiteDownloadStart float64 `json:"cost_per_onsite_download_start,string" gorm:"column:cost_per_onsite_download_start"`
	CostPerResult              float64 `json:"cost_per_result,string" gorm:"column:cost_per_result"`
	CostPerSecondaryGoalResult string  `json:"cost_per_secondary_goal_result" gorm:"column:cost_per_secondary_goal_result"`

	Cpc                       float64 `json:"cpc,string" gorm:"column:cpc"`
	Cpm                       float64 `json:"cpm,string" gorm:"column:cpm"`
	Ctr                       float64 `json:"ctr,string" gorm:"column:ctr"`
	Currency                  string  `json:"currency" gorm:"column:currency"`
	DpaTargetAudienceType     string  `json:"dpa_target_audience_type" gorm:"column:dpa_target_audience_type"`
	Frequency                 float64 `json:"frequency,string" gorm:"column:frequency"`
	GrossImpressions          uint64  `json:"gross_impressions,string" gorm:"column:gross_impressions"`
	Impressions               uint64  `json:"impressions,string" gorm:"column:impressions"`
	MobileAppId               uint64  `json:"mobile_app_id,string" gorm:"column:mobile_app_id"`
	ObjectiveType             string  `json:"objective_type" gorm:"column:objective_type"`
	OnOnsiteDownloadStart     uint64  `json:"onsite_download_start,string" gorm:"column:onsite_download_start"`
	OnsiteDownloadStartRate   float64 `json:"onsite_download_start_rate,string" gorm:"column:onsite_download_start_rate"`
	OptStatus                 int     `json:"opt_status,string" gorm:"column:opt_status"`
	PlacementType             string  `json:"placement_type" gorm:"column:placement_type"`
	PricingCategory           int     `json:"pricing_category,string" gorm:"column:pricing_category"`
	PromotionType             string  `json:"promotion_type" gorm:"column:promotion_type"`
	Reach                     uint64  `json:"reach,string" gorm:"column:reach"`
	RealTimeConversion        uint64  `json:"real_time_conversion,string" gorm:"column:real_time_conversion"`
	RealTimeConversionRate    float64 `json:"real_time_conversion_rate,string" gorm:"column:real_time_conversion_rate"`
	RealTimeCostPerConversion float64 `json:"real_time_cost_per_conversion,string" gorm:"column:real_time_cost_per_conversion"`
	RealTimeCostPerResult     float64 `json:"real_time_cost_per_result,string" gorm:"column:real_time_cost_per_result"`
	RealTimeResult            uint64  `json:"real_time_result,string" gorm:"column:real_time_result"`
	RealTimeResultRate        float64 `json:"real_time_result_rate,string" gorm:"column:real_time_result_rate"`
	Result                    int     `json:"result,string" gorm:"column:result"`
	ResultRate                float64 `json:"result_rate,string" gorm:"column:result_rate"`
	SecondaryGoalResult       string  `json:"secondary_goal_result" gorm:"column:secondary_goal_result"`
	SecondaryGoalResultRate   string  `json:"secondary_goal_result_rate" gorm:"column:secondary_goal_result_rate"`
	SmartTarget               string  `json:"smart_target" gorm:"column:smart_target"`
	Spend                     float64 `json:"spend,string" gorm:"column:spend"`
	SplitTest                 int     `json:"split_test,string" gorm:"column:split_test"`
	TtAppId                   uint64  `json:"tt_app_id,string" gorm:"column:tt_app_id"`
	TtAppName                 string  `json:"tt_app_name" gorm:"column:tt_app_name"`
	archived.NoPKModel
}

func (TiktokAdsAdGroupReport) TableName() string {
	return "_tool_tiktokAds_ad_group_report"
}
