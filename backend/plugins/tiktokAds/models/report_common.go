package models

import "github.com/apache/incubator-devlake/core/models/common"

type TiktokAdsReportCommon struct {
	StatTimeDay                string  `json:"stat_time_day" gorm:"column:stat_time_day;primaryKey" mapstructure:"stat_time_day"`
	AeoType                    string  `json:"aeo_type" gorm:"column:aeo_type" mapstructure:"aeo_type"`
	AppPromotionType           string  `json:"app_promotion_type" gorm:"column:app_promotion_type" mapstructure:"app_promotion_type"`
	Bid                        string  `json:"bid" gorm:"column:bid" mapstructure:"bid"`
	BidStrategy                string  `json:"bid_strategy" gorm:"column:bid_strategy" mapstructure:"bid_strategy"`
	Budget                     float64 `json:"budget,string" gorm:"column:budget" mapstructure:"budget"`
	CampaignBudget             float64 `json:"campaign_budget,string" gorm:"column:campaign_budget" mapstructure:"campaign_budget"`
	CampaignDedicateType       string  `json:"campaign_dedicate_type" gorm:"column:campaign_dedicate_type" mapstructure:"campaign_dedicate_type"`
	CampaignId                 uint64  `json:"campaign_id,string" gorm:"column:campaign_id" mapstructure:"campaign_id"`
	CampaignName               string  `json:"campaign_name" gorm:"column:campaign_name" mapstructure:"campaign_name"`
	Clicks                     int     `json:"clicks,string" gorm:"column:clicks" mapstructure:"clicks"`
	Conversion                 int     `json:"conversion,string" gorm:"column:conversion" mapstructure:"conversion"`
	ConversionRate             float64 `json:"conversion_rate,string" gorm:"column:conversion_rate" mapstructure:"conversion_rate"`
	CostPer1000Reached         float64 `json:"cost_per_1000_reached,string" gorm:"column:cost_per_1000_reached" mapstructure:"cost_per_1000_reached"`
	CostPerConversion          float64 `json:"cost_per_conversion,string" gorm:"column:cost_per_conversion" mapstructure:"cost_per_conversion"`
	CostPerOnsiteDownloadStart float64 `json:"cost_per_onsite_download_start,string" gorm:"column:cost_per_onsite_download_start" mapstructure:"cost_per_onsite_download_start"`
	CostPerResult              float64 `json:"cost_per_result,string" gorm:"column:cost_per_result" mapstructure:"cost_per_result"`
	CostPerSecondaryGoalResult string  `json:"cost_per_secondary_goal_result" gorm:"column:cost_per_secondary_goal_result" mapstructure:"cost_per_secondary_goal_result"`
	Cpc                        float64 `json:"cpc,string" gorm:"column:cpc" mapstructure:"cpc"`
	Cpm                        float64 `json:"cpm,string" gorm:"column:cpm" mapstructure:"cpm"`
	Ctr                        float64 `json:"ctr,string" gorm:"column:ctr" mapstructure:"ctr"`
	Currency                   string  `json:"currency" gorm:"column:currency" mapstructure:"currency"`
	DpaTargetAudienceType      string  `json:"dpa_target_audience_type" gorm:"column:dpa_target_audience_type" mapstructure:"dpa_target_audience_type"`
	Frequency                  float64 `json:"frequency,string" gorm:"column:frequency" mapstructure:"frequency"`
	GrossImpressions           uint64  `json:"gross_impressions,string" gorm:"column:gross_impressions" mapstructure:"gross_impressions"`
	Impressions                uint64  `json:"impressions,string" gorm:"column:impressions" mapstructure:"impressions"`
	MobileAppId                uint64  `json:"mobile_app_id,string" gorm:"column:mobile_app_id" mapstructure:"mobile_app_id"`
	ObjectiveType              string  `json:"objective_type" gorm:"column:objective_type" mapstructure:"objective_type"`
	OnOnsiteDownloadStart      uint64  `json:"onsite_download_start,string" gorm:"column:onsite_download_start" mapstructure:"onsite_download_start"`
	OnsiteDownloadStartRate    float64 `json:"onsite_download_start_rate,string" gorm:"column:onsite_download_start_rate" mapstructure:"onsite_download_start_rate"`
	OptStatus                  int     `json:"opt_status,string" gorm:"column:opt_status" mapstructure:"opt_status"`
	PlacementType              string  `json:"placement_type" gorm:"column:placement_type" mapstructure:"placement_type"`
	PricingCategory            int     `json:"pricing_category,string" gorm:"column:pricing_category" mapstructure:"pricing_category"`
	PromotionType              string  `json:"promotion_type" gorm:"column:promotion_type" mapstructure:"promotion_type"`
	Reach                      uint64  `json:"reach,string" gorm:"column:reach" mapstructure:"reach"`
	RealTimeConversion         uint64  `json:"real_time_conversion,string" gorm:"column:real_time_conversion" mapstructure:"real_time_conversion"`
	RealTimeConversionRate     float64 `json:"real_time_conversion_rate,string" gorm:"column:real_time_conversion_rate" mapstructure:"real_time_conversion_rate"`
	RealTimeCostPerConversion  float64 `json:"real_time_cost_per_conversion,string" gorm:"column:real_time_cost_per_conversion" mapstructure:"real_time_cost_per_conversion"`
	RealTimeCostPerResult      float64 `json:"real_time_cost_per_result,string" gorm:"column:real_time_cost_per_result" mapstructure:"real_time_cost_per_result"`
	RealTimeResult             uint64  `json:"real_time_result,string" gorm:"column:real_time_result" mapstructure:"real_time_result"`
	RealTimeResultRate         float64 `json:"real_time_result_rate,string" gorm:"column:real_time_result_rate" mapstructure:"real_time_result_rate"`
	Result                     int     `json:"result,string" gorm:"column:result" mapstructure:"result"`
	ResultRate                 float64 `json:"result_rate,string" gorm:"column:result_rate" mapstructure:"result_rate"`
	SecondaryGoalResult        string  `json:"secondary_goal_result" gorm:"column:secondary_goal_result" mapstructure:"secondary_goal_result"`
	SecondaryGoalResultRate    string  `json:"secondary_goal_result_rate" gorm:"column:secondary_goal_result_rate" mapstructure:"secondary_goal_result_rate"`
	SmartTarget                string  `json:"smart_target" gorm:"column:smart_target" mapstructure:"smart_target"`
	Spend                      float64 `json:"spend,string" gorm:"column:spend" mapstructure:"spend"`
	SplitTest                  int     `json:"split_test,string" gorm:"column:split_test" mapstructure:"split_test"`
	TtAppId                    uint64  `json:"tt_app_id,string" gorm:"column:tt_app_id" mapstructure:"tt_app_id"`
	TtAppName                  string  `json:"tt_app_name" gorm:"column:tt_app_name" mapstructure:"tt_app_name"`
	common.NoPKModel
}
