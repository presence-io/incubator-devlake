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

import (
	"github.com/apache/incubator-devlake/core/models/common"
)

type TiktokAdsModifyHistory struct {
	ConnectionId uint64 `json:"connection_id" gorm:"column:connection_id;primaryKey;autoIncrement:false"`
	AdvertiserID string `gorm:"column:advertiser_id" json:"advertiser_id"`
	StatTimeDay  string `gorm:"column:stat_time_day;index" json:"stat_time_day"`
	AdgroupId    uint64 `gorm:"column:adgroup_id;index" json:"adgroup_id"`
	AdId         uint64 `gorm:"column:ad_id;index" json:"ad_id"`
	ModifyField  string `gorm:"column:modify_field;index" json:"modify_field"`
	CurrentValue string `gorm:"column:current_value" json:"current_value"`
	common.NoPKModel
}

func (TiktokAdsModifyHistory) TableName() string {
	return "_tool_tiktokAds_modify_history"
}
