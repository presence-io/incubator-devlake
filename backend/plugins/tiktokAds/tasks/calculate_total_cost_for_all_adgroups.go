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
	"gorm.io/gorm"
	"time"
)

var _ plugin.SubTaskEntryPoint = CalculateTotalCostForAllAdgroups

func CalculateTotalCostForAllAdgroups(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	db := taskCtx.GetDal()
	notify := &models.TiktokAdsNotifyHistory{}

	err := db.First(notify, dal.Where("connection_id = ? and stat_time_day = ? and advertiser_id = ?",
		data.Options.ConnectionId, time.Now().Format("2006-01-02 00:00:00"), data.Options.AdvertiserID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if notify.IsNotifyBudget {
		return nil
	}

	adGroupReportClauses := []dal.Clause{
		dal.Select("sum(spend) as total_cost"),
		dal.From(&models.TiktokAdsAdGroupReport{}),
	}
	adGroupReportClauses = append(adGroupReportClauses, dal.Where("connection_id = ? and stat_time_day = ? and adgroup_id in (?)",
		data.Options.ConnectionId, time.Now().Format("2006-01-02 00:00:00"), data.Options.AdGroupIds))

	total := 0.0
	err = db.First(&total, adGroupReportClauses...)
	if err != nil {
		return errors.Convert(err)
	}
	if total >= 500 {
		resMsg := fmt.Sprintf("总花费超过500预警！！！: %f; <at user_id=\"all\">所有人</at>", total)
		feishuNotify(resMsg, nil, data.ApiClient)
		notify.IsNotifyBudget = true
		notify.AdvertiserID = data.Options.AdvertiserID
		notify.StatTimeDay = time.Now().Format("2006-01-02 00:00:00")
		notify.ConnectionId = data.Options.ConnectionId
		err = db.Create(notify)
	}
	if err != nil {
		return err
	}

	return nil
}

var CalculateTotalCostForAllAdgroupsMeta = plugin.SubTaskMeta{
	Name:             "CalculateTotalCostForAllAdgroups",
	EntryPoint:       CalculateTotalCostForAllAdgroups,
	EnabledByDefault: true,
	Description:      "",
	DomainTypes:      []string{},
}
