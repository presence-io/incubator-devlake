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
	"github.com/apache/incubator-devlake/plugins/tiktokAds/api"
	"github.com/apache/incubator-devlake/plugins/tiktokAds/models"
	"github.com/mitchellh/mapstructure"
	"strings"
	"time"
)

var _ plugin.SubTaskEntryPoint = ExecuteAdGroupRules

func ExecuteAdGroupRules(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	db := taskCtx.GetDal()
	logger := taskCtx.GetLogger()
	now := time.Now() // 获取当前时间
	rules := make([]*models.TiktokAdsRule, 0)
	clauses := []dal.Clause{
		dal.From(&models.TiktokAdsRule{}),
		dal.Where("connection_id = ? and data_level = ? and status = ? and id in (?)",
			data.Options.ConnectionId, "adgroup", models.Active, data.Options.RuleIds),
	}
	err := db.All(&rules, clauses...)
	if err != nil {
		return errors.Convert(err)
	}
	ruleConditionMap := make(map[models.TiktokAdsRule][]*models.TiktokAdsRuleCondition)
	enableAdGroup := make([]uint64, 0)
	disableAdGroup := make([]uint64, 0)
	reviseAdGroupModify := make(map[models.TiktokAdsRule][]models.TiktokAdsAdGroupReport)
	adGroupReports := make([]*models.TiktokAdsAdGroupReport, 0)

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

		adGroupReportClauses := []dal.Clause{
			dal.From(&models.TiktokAdsAdGroupReport{}),
		}
		if len(rule.AdGroupIds) > 0 && rule.AdGroupIds != "[]" {
			adGroupIds := strings.Split(rule.AdGroupIds, "|")
			adGroupReportClauses = append(adGroupReportClauses, dal.Where("connection_id = ? and stat_time_day = ? and adgroup_id in (?)", data.Options.ConnectionId, now.Format("2006-01-02 00:00:00"), adGroupIds))
		} else {
			adGroupReportClauses = append(adGroupReportClauses, dal.Where("connection_id = ? and stat_time_day = ?", data.Options.ConnectionId, now.Format("2006-01-02 00:00:00")))
		}
		err = db.All(&adGroupReports, adGroupReportClauses...)
		if err != nil {
			return errors.Convert(err)
		}
		for _, adGroupReport := range adGroupReports {

			modifyHistoryClauses := []dal.Clause{
				dal.From(&models.TiktokAdsModifyHistory{}),
				dal.Where("connection_id = ? and stat_time_day = ? and adgroup_id = ? and modify_field in (?)",
					data.Options.ConnectionId, now.Format("2006-01-02 00:00:00"), adGroupReport.AdgroupId, []string{"conversion_bid_price"}),
			}
			count, err := db.Count(modifyHistoryClauses...)
			if err != nil {
				return err
			}
			if count >= api.MaxModifyCount {
				continue
			}
			logger.Info(fmt.Sprintf("开始执行 adGroupReport: %s , 匹配ruleId %d ", adGroupReport.AdgroupName, rule.ID))
			if adGroupReport.OptStatus == models.Delete {
				continue
			}
			// convert adReport to a map
			reportValueMap := make(map[string]interface{})
			err = errors.Convert(mapstructure.Decode(&adGroupReport.TiktokAdsReportCommon, &reportValueMap))
			if err != nil {
				return errors.Convert(err)
			}
			reportValueMap["id"] = adGroupReport.AdgroupId
			if rule.Operate == models.ENABLE && adGroupReport.OptStatus == models.Active {
				continue
			}
			if rule.Operate == models.DISABLE && adGroupReport.OptStatus == models.InActive {
				continue
			}
			// 如果bid strategy 不是 Cost Cap，就不能调整出价
			if rule.Operate == models.MODIFY && rule.FieldToRevise == "conversion_bid_price" && adGroupReport.BidStrategy != models.COST_CAP {
				continue
			}
			if calculateRule(conditions, reportValueMap, taskCtx) {
				logger.Info(fmt.Sprintf("adGroupReport: %s , 已经满足ruleId %d ，即将进行下一步操作 <%s>; 需要调整的字段<%s>; 原始值<%.2f>; 调整比例<%.2f>",
					adGroupReport.AdgroupName, rule.ID, rule.Operate, rule.FieldToRevise, reportValueMap[rule.FieldToRevise], rule.ValueToRevise))
				switch rule.Operate {
				case models.ENABLE:
					enableAdGroup = append(enableAdGroup, adGroupReport.AdgroupId)
				case models.DISABLE:
					disableAdGroup = append(disableAdGroup, adGroupReport.AdgroupId)
				case models.MODIFY:
					reviseAdGroupModify[*rule] = append(reviseAdGroupModify[*rule], *adGroupReport)
				}
			}
		}
	}
	if len(enableAdGroup) > 0 {
		prepareUpdate("adgroup", enableAdGroup, taskCtx, models.ENABLE)
	}
	if len(disableAdGroup) > 0 {
		prepareUpdate("adgroup", disableAdGroup, taskCtx, models.DISABLE)
	}
	if len(reviseAdGroupModify) > 0 {
		modifyField(reviseAdGroupModify, taskCtx)
	}

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
