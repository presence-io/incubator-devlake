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

package models

type TiktokAdsAdReport struct {
	ConnectionId uint64 `json:"connection_id" gorm:"column:connection_id;autoIncrement:false;primaryKey" mapstructure:"connection_id"`
	StatTimeDay  string `json:"stat_time_day" gorm:"column:stat_time_day;primaryKey" mapstructure:"stat_time_day"`
	AdId         uint64 `json:"ad_id,string" gorm:"column:ad_id;primaryKey;autoIncrement:false" mapstructure:"ad_id"`
	AdvertiserID string `json:"advertiser_id,omitempty" gorm:"column:advertiser_id;primaryKey" mapstructure:"advertiser_id,omitempty"`
	AdGroupId    uint64 `json:"adgroup_id,string" gorm:"column:adgroup_id" mapstructure:"adgroup_id"`
	AdGroupName  string `json:"adgroup_name" gorm:"column:adgroup_name" mapstructure:"adgroup_name"`
	AdName       string `json:"ad_name" gorm:"column:ad_name" mapstructure:"ad_name"`
	AdText       string `json:"ad_text" gorm:"column:ad_text" mapstructure:"ad_text"`
	CallToAction string `json:"call_to_action" gorm:"column:call_to_action" mapstructure:"call_to_action"`
	ImageMode    string `json:"image_mode" gorm:"column:image_mode" mapstructure:"image_mode"`
	TiktokAdsReportCommon
}

func (TiktokAdsAdReport) TableName() string {
	return "_tool_tiktokAds_ad_report"
}

type AdReportListItem struct {
	Metrics    TiktokAdsAdReport `json:"metrics"`
	Dimensions Dimensions        `json:"dimensions"`
}

type AdReportResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		PageInfo PageInfo           `json:"page_info"`
		List     []AdReportListItem `json:"list"`
	} `json:"data"`
	RequestID string `json:"request_id"`
}
