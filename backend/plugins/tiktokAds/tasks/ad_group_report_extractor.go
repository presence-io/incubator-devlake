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
	"encoding/json"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/tiktokAds/models"
)

var _ plugin.SubTaskEntryPoint = ExtractAdGroupReport

func ExtractAdGroupReport(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TiktokAdsApiParams{
				ConnectionId: data.Options.ConnectionId,
				AdvertiserId: data.Options.AdvertiserID,
			},
			Table: RAW_AD_GROUP_REPORT_TABLE,
		},
		Extract: func(resData *helper.RawData) ([]interface{}, errors.Error) {
			extractedModels := make([]interface{}, 0)
			res := &models.AdGroupReportListItem{}
			err := json.Unmarshal(resData.Data, res)
			if err != nil {
				return nil, errors.Convert(err)
			}
			report := res.Metrics
			report.AdgroupId = res.Dimensions.AdGroupID
			report.ConnectionId = data.Options.ConnectionId
			report.StatTimeDay = res.Dimensions.StatTimeDay
			report.AdvertiserID = data.Options.AdvertiserID
			extractedModels = append(extractedModels, &report)
			return extractedModels, nil
		},
	})
	if err != nil {
		return err
	}

	return extractor.Execute()
}

var ExtractAdGroupReportMeta = plugin.SubTaskMeta{
	Name:             "ExtractAdGroupReport",
	EntryPoint:       ExtractAdGroupReport,
	EnabledByDefault: true,
	Description:      "Extract raw data into tool layer table tiktokads_report",
	DomainTypes:      []string{},
}
