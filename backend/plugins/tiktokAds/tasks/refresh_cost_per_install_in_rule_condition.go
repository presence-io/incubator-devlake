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
	"time"
)

var _ plugin.SubTaskEntryPoint = RefreshCostPerInstall

func RefreshCostPerInstall(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	db := taskCtx.GetDal()

	adGroupReportClauses := []dal.Clause{
		dal.Select("sum(spend) / sum(onsite_download_start) as avg_cost"),
		dal.From(&models.TiktokAdsAdGroupReport{}),
	}
	adGroupReportClauses = append(adGroupReportClauses, dal.Where("connection_id = ? and stat_time_day = ? and adgroup_id in (?)",
		data.Options.ConnectionId, time.Now().Format("2006-01-02 00:00:00"), data.Options.AdGroupIds))

	avg := 0.0
	err := db.First(&avg, adGroupReportClauses...)
	if err != nil {
		return errors.Convert(err)
	}

	updateClauses := []dal.Clause{
		dal.Where("field_name = ? and rule_id in (?)",
			"cost_per_onsite_download_start", []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}),
	}
	err = db.UpdateColumn(&models.TiktokAdsRuleCondition{}, "field_value", fmt.Sprintf("%.2f", avg+0.02), updateClauses...)
	if err != nil {
		return err
	}

	return nil
}

var RefreshCostPerInstallMeta = plugin.SubTaskMeta{
	Name:             "RefreshCostPerInstall",
	EntryPoint:       RefreshCostPerInstall,
	EnabledByDefault: true,
	Description:      "",
	DomainTypes:      []string{},
}
