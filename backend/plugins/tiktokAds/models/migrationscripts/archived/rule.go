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

import (
	"github.com/apache/incubator-devlake/core/models/migrationscripts/archived"
)

type TiktokAdsRule struct {
	archived.Model
	ConnectionId  uint64 `json:"connection_id" gorm:"column:connection_id;autoIncrement:false;primaryKey"`
	Name          string `json:"name" gorm:"column:name"`
	CampaignId    uint64 `json:"campaign_id,string" gorm:"column:campaign_id"`
	AdgroupId     uint64 `json:"adgroup_id" gorm:"column:adgroup_id"`
	AdId          uint64 `json:"ad_id" gorm:"column:ad_id"`
	Status        int
	Operate       string
	FieldToRevise string  `json:"field_to_revise" mapstructure:"field_to_revise"`
	ValueToRevise float64 `json:"value_to_revise" mapstructure:"value_to_revise"`
	DataLevel     string
	AdGroupIds    string `json:"adgroup_ids" mapstructure:"adgroup_ids"`
	AdIds         string `json:"ad_ids" mapstructure:"ad_ids"`
}

type TiktokAdsRuleCondition struct {
	ID                uint64 `gorm:"primaryKey"`
	RuleID            uint64
	FieldName         string
	FieldRelationName string  `json:"field_relation_name" mapstructure:"field_relation_name"`
	Operator          string  `json:"operator" mapstructure:"operator"`
	FieldValue        float64 `json:"field_value" mapstructure:"field_value"`
	FieldRelationRate float64 `json:"field_relation_rate" mapstructure:"field_relation_rate"`
	TimeRange         uint64  //hours
	IsEnable          bool
	archived.Model
}

func (TiktokAdsRule) TableName() string {
	return "_tool_tiktokAds_rule"
}

func (TiktokAdsRuleCondition) TableName() string {
	return "_tool_tiktokAds_rule_condition"
}
