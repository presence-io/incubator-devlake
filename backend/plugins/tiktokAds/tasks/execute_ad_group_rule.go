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

var _ plugin.SubTaskEntryPoint = ExecuteAdGroupRules

func ExecuteAdGroupRules(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	db := taskCtx.GetDal()
	now := time.Now() // 获取当前时间
	rules := make([]*models.TiktokAdsRule, 0)
	clauses := []dal.Clause{
		dal.From(&models.TiktokAdsRule{}),
		dal.Where("connection_id = ? and data_level = ? ", data.Options.ConnectionId, "ad_group"),
	}
	err := db.All(&rules, clauses...)
	if err != nil {
		return errors.Convert(err)
	}
	ruleConditionMap := make(map[models.TiktokAdsRule][]*models.TiktokAdsRuleCondition)
	enableAdGroup := make([]uint64, 0)
	disableAdGroup := make([]uint64, 0)
	reviseAdGroupBudget := make(map[models.TiktokAdsRule]models.TiktokAdsAdGroupReport)

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

	adGroupReports := make([]*models.TiktokAdsAdGroupReport, 0)
	adGroupReportClauses := []dal.Clause{
		dal.From(&models.TiktokAdsAdGroupReport{}),
		dal.Where("connection_id = ? and stat_time_day = ?", data.Options.ConnectionId, now.Format("2006-01-02")),
	}
	err = db.All(&adGroupReports, adGroupReportClauses...)
	if err != nil {
		return errors.Convert(err)
	}
	for _, adGroupReport := range adGroupReports {
		if adGroupReport.OptStatus == models.Delete {
			continue
		}
		// convert adReport to a map
		reportValueMap := make(map[string]interface{})
		err = errors.Convert(mapstructure.Decode(&adGroupReport.TiktokAdsReportCommon, reportValueMap))
		if err != nil {
			return errors.Convert(err)
		}
		for rule, conditions := range ruleConditionMap {
			if rule.Operate == models.ENABLE && adGroupReport.OptStatus == models.Active {
				continue
			}
			if rule.Operate == models.DISABLE && adGroupReport.OptStatus == models.InActive {
				continue
			}
			if calculateRule(rule, conditions, reportValueMap) {
				switch rule.Operate {
				case models.ENABLE:
					enableAdGroup = append(enableAdGroup, adGroupReport.AdgroupId)
				case models.DISABLE:
					disableAdGroup = append(disableAdGroup, adGroupReport.AdgroupId)
				case models.MODIFY_BUDGET:
					reviseAdGroupBudget[rule] = *adGroupReport
				}

			}
		}
	}

	prepareUpdate("adgroup", enableAdGroup, data, models.ENABLE)
	prepareUpdate("adgroup", disableAdGroup, data, models.DISABLE)
	modifyBudget(reviseAdGroupBudget, data)

	if err != nil {
		return errors.Convert(err)
	}
	return nil
}

var ExecuteAdGroupRulesMeta = plugin.SubTaskMeta{
	Name:             "ExecuteAdGroupRules",
	EntryPoint:       ExecuteAdGroupRules,
	EnabledByDefault: true,
	Description:      "",
	DomainTypes:      []string{},
}
