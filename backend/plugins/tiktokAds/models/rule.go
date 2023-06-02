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

type TiktokAdsRule struct {
	ConnectionId uint64 `json:"connection_id" gorm:"column:connection_id;primaryKey;autoIncrement:false" mapstructure:"connection_id"`
	common.Model
	Name          string  `json:"name" gorm:"column:name" mapstructure:"name"`
	CampaignId    uint64  `json:"campaign_id" gorm:"column:campaign_id" mapstructure:"campaign_id"`
	AdgroupId     uint64  `json:"adgroup_id" gorm:"column:adgroup_id" mapstructure:"adgroup_id"`
	AdId          uint64  `json:"ad_id" gorm:"column:ad_id" mapstructure:"ad_id"`
	Status        int     `json:"status" mapstructure:"status"`
	Operate       string  `json:"operate" mapstructure:"operate"`
	FieldToRevise string  `json:"field_to_revise" mapstructure:"field_to_revise"`
	ValueToRevise float64 `json:"value_to_revise" mapstructure:"value_to_revise"`
	DataLevel     string  `json:"data_level" mapstructure:"data_level"`
	AdGroupIds    string  `json:"adgroup_ids" mapstructure:"adgroup_ids"`
	AdIds         string  `json:"ad_ids" mapstructure:"ad_ids"`
}

type TiktokAdsRuleCondition struct {
	RuleID            uint64  `json:"rule_id" mapstructure:"rule_id"`
	FieldName         string  `json:"field_name" mapstructure:"field_name"`
	FieldRelationName string  `json:"field_relation_name" mapstructure:"field_relation_name"`
	Operator          string  `json:"operator" mapstructure:"operator"`
	FieldValue        float64 `json:"field_value" mapstructure:"field_value"`
	FieldRelationRate float64 `json:"field_relation_rate" mapstructure:"field_relation_rate"`
	TimeRange         uint64  `json:"time_range" mapstructure:"time_range"`
	IsEnable          bool    `json:"is_enable" mapstructure:"is_enable"`
	common.Model
}

type TiktokAdsRuleRequest struct {
	ConnectionId  uint64                    `json:"connection_id" gorm:"column:connection_id;primaryKey;autoIncrement:false" mapstructure:"connection_id"`
	ID            uint64                    `json:"id" gorm:"column:id;primaryKey;autoIncrement:false" mapstructure:"id"`
	Name          string                    `json:"name" gorm:"column:name" mapstructure:"name"`
	CampaignId    uint64                    `json:"campaign_id" gorm:"column:campaign_id" mapstructure:"campaign_id"`
	AdgroupId     uint64                    `json:"adgroup_id" gorm:"column:adgroup_id" mapstructure:"adgroup_id"`
	AdId          uint64                    `json:"ad_id" gorm:"column:ad_id"mapstructure:"ad_id"`
	Status        int                       `json:"status" mapstructure:"status"`
	Operate       string                    `json:"operate" mapstructure:"operate"`
	FieldToRevise string                    `json:"field_to_revise" mapstructure:"field_to_revise"`
	ValueToRevise float64                   `json:"value_to_revise" mapstructure:"value_to_revise"`
	DataLevel     string                    `json:"data_level" mapstructure:"data_level"`
	AdGroupIds    []string                  `json:"adgroup_ids" mapstructure:"adgroup_ids"`
	AdIds         []string                  `json:"ad_ids" mapstructure:"ad_ids"`
	Conditions    []*TiktokAdsRuleCondition `json:"conditions" mapstructure:"conditions"`
}

const (
	MODIFY  = "MODIFY"
	ENABLE  = "ENABLE"
	DISABLE = "DISABLE"
)

const (
	Active = iota
	InActive
	Delete
)

const (
	COST_CAP    = "Cost Cap"
	LOWEST_COST = "Lowest cost"
)

func (TiktokAdsRule) TableName() string {
	return "_tool_tiktokAds_rule"
}

func (TiktokAdsRuleCondition) TableName() string {
	return "_tool_tiktokAds_rule_condition"
}
