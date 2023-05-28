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
	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/plugins/tiktokAds/models"
	"github.com/mitchellh/mapstructure"
	"time"
)

var _ plugin.SubTaskEntryPoint = ExecuteAdRules

func ExecuteAdRules(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	db := taskCtx.GetDal()
	now := time.Now() // 获取当前时间
	rules := make([]*models.TiktokAdsRule, 0)
	clauses := []dal.Clause{
		dal.From(&models.TiktokAdsRule{}),
		dal.Where("connection_id = ? and data_level = ? ", data.Options.ConnectionId, "ad"),
	}
	err := db.All(&rules, clauses...)
	if err != nil {
		return errors.Convert(err)
	}
	ruleConditionMap := make(map[models.TiktokAdsRule][]*models.TiktokAdsRuleCondition)
	enableAd := make([]uint64, 0)
	disableAd := make([]uint64, 0)
	for _, rule := range rules {
		conditionClauses := []dal.Clause{
			dal.From(&models.TiktokAdsRuleCondition{}),
			dal.Where("rule_id = ?", rule.ID),
		}
		conditions := make([]*models.TiktokAdsRuleCondition, 0)
		err = db.All(&conditions, conditionClauses...)
		if err != nil {
			return errors.Convert(err)
		}
		ruleConditionMap[*rule] = conditions
	}

	adReports := make([]*models.TiktokAdsAdReport, 0)
	adReportClauses := []dal.Clause{
		dal.From(&models.TiktokAdsAdReport{}),
		dal.Where("connection_id = ? and stat_time_day = ?", data.Options.ConnectionId, now.Format("2006-01-02")),
	}
	err = db.All(&adReports, adReportClauses...)
	if err != nil {
		return errors.Convert(err)
	}
	for _, adReport := range adReports {
		if adReport.OptStatus == models.Delete {
			continue
		}
		// convert adReport to a map
		adReportValueMap := make(map[string]interface{})
		err = errors.Convert(mapstructure.Decode(&adReport.TiktokAdsReportCommon, adReportValueMap))
		if err != nil {
			return errors.Convert(err)
		}
		for rule, conditions := range ruleConditionMap {
			if rule.Operate == models.ENABLE && adReport.OptStatus == models.Active {
				continue
			}
			if rule.Operate == models.DISABLE && adReport.OptStatus == models.InActive {
				continue
			}
			if calculateRule(rule, conditions, adReportValueMap) {
				switch rule.Operate {
				case models.ENABLE:
					enableAd = append(enableAd, adReport.AdId)
				case models.DISABLE:
					disableAd = append(disableAd, adReport.AdId)
				}

			}
		}
	}
	prepareUpdate("ad", enableAd, data, models.ENABLE)
	prepareUpdate("ad", disableAd, data, models.DISABLE)
	if err != nil {
		return errors.Convert(err)
	}
	return nil
}

var ExecuteAdRulesMeta = plugin.SubTaskMeta{
	Name:             "ExecuteAdRules",
	EntryPoint:       ExecuteAdRules,
	EnabledByDefault: true,
	Description:      "",
	DomainTypes:      []string{},
}

func calculateRule(rule models.TiktokAdsRule, conditions []*models.TiktokAdsRuleCondition, reportValueMap map[string]interface{}) bool {
	for _, condition := range conditions {
		var float64Value float64
		// check if adReportMap[condition.FieldName] can be converted to float64
		if _, ok := reportValueMap[condition.FieldName].(float64); !ok {
			// 需要确认是否需要继续
		} else {
			float64Value = reportValueMap[condition.FieldName].(float64)
		}
		switch condition.Operator {
		case ">":
			if float64Value > condition.FieldValue {
				continue
			} else {
				return false
			}
		case "<":
			if float64Value < condition.FieldValue {
				continue
			} else {
				return false
			}
		case "=":
			if float64Value == condition.FieldValue {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
