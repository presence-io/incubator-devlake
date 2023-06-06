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
	"fmt"
	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/plugins/tiktokAds/models"
	"github.com/mitchellh/mapstructure"
	"strings"
	"time"
)

var _ plugin.SubTaskEntryPoint = ExecuteAdRules

func ExecuteAdRules(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	db := taskCtx.GetDal()
	logger := taskCtx.GetLogger()
	now := time.Now() // 获取当前时间
	rules := make([]*models.TiktokAdsRule, 0)
	clauses := []dal.Clause{
		dal.From(&models.TiktokAdsRule{}),
		dal.Where("connection_id = ? and data_level = ? and status = ? and id in (?)",
			data.Options.ConnectionId, "ad", models.Active, data.Options.RuleIds),
	}
	err := db.All(&rules, clauses...)
	if err != nil {
		return errors.Convert(err)
	}
	ruleConditionMap := make(map[models.TiktokAdsRule][]*models.TiktokAdsRuleCondition)
	enableAd := make([]uint64, 0)
	disableAd := make([]uint64, 0)
	adReports := make([]*models.TiktokAdsAdReport, 0)

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

		adReportClauses := []dal.Clause{
			dal.From(&models.TiktokAdsAdReport{}),
		}
		if len(rule.AdIds) > 0 && rule.AdIds != "[]" {
			adIds := strings.Split(rule.AdIds, "|")
			adReportClauses = append(adReportClauses, dal.Where("connection_id = ? and stat_time_day = ? and ad_id in (?)", data.Options.ConnectionId, now.Format("2006-01-02 00:00:00"), adIds))
		} else if len(rule.AdGroupIds) > 0 && rule.AdGroupIds != "[]" {
			adGroupIds := strings.Split(rule.AdGroupIds, "|")
			adReportClauses = append(adReportClauses, dal.Where("connection_id = ? and stat_time_day = ? and adgroup_id in (?)", data.Options.ConnectionId, now.Format("2006-01-02 00:00:00"), adGroupIds))
		} else {
			adReportClauses = append(adReportClauses, dal.Where("connection_id = ? and stat_time_day = ?", data.Options.ConnectionId, now.Format("2006-01-02 00:00:00")))
		}
		err = db.All(&adReports, adReportClauses...)
		if err != nil {
			return errors.Convert(err)
		}
		for _, adReport := range adReports {
			logger.Info(fmt.Sprintf("开始执行 ad : %s , 匹配ruleId %d ", adReport.AdName, rule.ID))
			if adReport.OptStatus == models.Delete {
				continue
			}
			// convert adReport to a map
			reportValueMap := make(map[string]interface{})
			err = errors.Convert(mapstructure.Decode(&adReport.TiktokAdsReportCommon, &reportValueMap))
			if err != nil {
				return errors.Convert(err)
			}
			if rule.Operate == models.ENABLE && adReport.OptStatus == models.Active {
				continue
			}
			if rule.Operate == models.DISABLE && adReport.OptStatus == models.InActive {
				continue
			}
			// 如果bid strategy 不是 Cost Cap，就不能调整出价
			if calculateRule(conditions, reportValueMap, taskCtx) {
				if rule.Operate == models.ENABLE && adReport.OptStatus == models.Active {
					continue
				}
				if rule.Operate == models.DISABLE && adReport.OptStatus == models.InActive {
					continue
				}
				logger.Info(fmt.Sprintf("adReport: %s , 满足ruleId %d 进行下一步操作", adReport.AdName, rule.ID))
				switch rule.Operate {
				case models.ENABLE:
					enableAd = append(enableAd, adReport.AdId)
				case models.DISABLE:
					disableAd = append(disableAd, adReport.AdId)
				}
			}
		}
	}
	if len(enableAd) > 0 {
		prepareUpdate("ad", enableAd, taskCtx, models.ENABLE)
	}
	if len(disableAd) > 0 {
		prepareUpdate("ad", disableAd, taskCtx, models.DISABLE)
	}
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
