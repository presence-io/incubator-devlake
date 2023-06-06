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
	"time"
)

var _ plugin.SubTaskEntryPoint = ExtractReport

func ExtractReport(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TiktokAdsApiParams{
				ConnectionId: data.Options.ConnectionId,
				AdvertiserId: data.Options.AdvertiserID,
				StatTimeDate: time.Now().Format("2006-01-02 00:00:00"),
			},
			Table: RAW_REPORT_TABLE,
		},
		Extract: func(resData *helper.RawData) ([]interface{}, errors.Error) {
			extractedModels := make([]interface{}, 0)
			// TODO decode report models from api result
			res := &models.ReportListItem{}
			err := json.Unmarshal(resData.Data, res)
			if err != nil {
				return nil, errors.Convert(err)
			}
			report := &models.TiktokAdsReport{
				//ConnectionId:               data.Options.ConnectionId,
				//AdvertiserID:               res.Dimensions.AdvertiserID,
				//CampaignID:                 res.Dimensions.CampaignID,
				//AdGroupID:                  res.Dimensions.AdGroupID,
				//AdID:                       res.Dimensions.AdID,
				//StatTimeDay:                res.Dimensions.StatTimeDay,
				//StatTimeHour:               res.Dimensions.StatTimeHour,
				//AC:                         res.Dimensions.AC,
				//Age:                        res.Dimensions.Age,
				//CountryCode:                res.Dimensions.CountryCode,
				//InterestCategory:           res.Dimensions.InterestCategory,
				//InterestCategoryV2:         res.Dimensions.InterestCategoryV2,
				//Gender:                     res.Dimensions.Gender,
				//Language:                   res.Dimensions.Language,
				//Placement:                  res.Dimensions.Placement,
				//Platform:                   res.Dimensions.Platform,
				//ContextualTag:              res.Dimensions.ContextualTag,
				//Spend:                      res.Metrics.Spend,
				//CashSpend:                  res.Metrics.CashSpend,
				//VoucherSpend:               res.Metrics.VoucherSpend,
				//CPC:                        res.Metrics.CPC,
				//CPM:                        res.Metrics.CPM,
				//Impressions:                res.Metrics.Impressions,
				//GrossImpressions:           res.Metrics.GrossImpressions,
				//Clicks:                     res.Metrics.Clicks,
				//CTR:                        res.Metrics.CTR,
				//Reach:                      res.Metrics.Reach,
				//CostPer1000Reached:         res.Metrics.CostPer1000Reached,
				//Conversion:                 res.Metrics.Conversion,
				//CostPerConversion:          res.Metrics.CostPerConversion,
				//ConversionRate:             res.Metrics.ConversionRate,
				//RealTimeConversion:         res.Metrics.RealTimeConversion,
				//RealTimeCostPerConversion:  res.Metrics.RealTimeCostPerConversion,
				//RealTimeConversionRate:     res.Metrics.RealTimeConversionRate,
				//Result:                     res.Metrics.Result,
				//CostPerResult:              res.Metrics.CostPerResult,
				//ResultRate:                 res.Metrics.ResultRate,
				//RealTimeResult:             res.Metrics.RealTimeResult,
				//RealTimeCostPerResult:      res.Metrics.RealTimeCostPerResult,
				//RealTimeResultRate:         res.Metrics.RealTimeResultRate,
				//SecondaryGoalResult:        res.Metrics.SecondaryGoalResult,
				//CostPerSecondaryGoalResult: res.Metrics.CostPerSecondaryGoalResult,
				//SecondaryGoalResultRate:    res.Metrics.SecondaryGoalResultRate,
				//Frequency:                  res.Metrics.Frequency,
				//Currency:                   res.Metrics.Currency,
				//CampaignName:               res.Metrics.CampaignName,
				//ObjectiveType:              res.Metrics.ObjectiveType,
				//SplitTest:                  res.Metrics.SplitTest,
				//CampaignBudget:             res.Metrics.CampaignBudget,
				//CampaignDedicateType:       res.Metrics.CampaignDedicateType,
				//AppPromotionType:           res.Metrics.AppPromotionType,
				//AdGroupName:                res.Metrics.AdGroupName,
				//PlacementType:              res.Metrics.PlacementType,
				//PromotionType:              res.Metrics.PromotionType,
				//OptStatus:                  res.Metrics.OptStatus,
				//DpaTargetAudienceType:      res.Metrics.DpaTargetAudienceType,
				//Budget:                     res.Metrics.Budget,
				//SmartTarget:                res.Metrics.SmartTarget,
				//PricingCategory:            res.Metrics.PricingCategory,
				//BidStrategy:                res.Metrics.BidStrategy,
				//Bid:                        res.Metrics.Bid,
				//AeoType:                    res.Metrics.AeoType,
				//AdName:                     res.Metrics.AdName,
				//AdText:                     res.Metrics.AdText,
				//CallToAction:               res.Metrics.CallToAction,
				//TikTokAppID:                res.Metrics.TikTokAppID,
				//TikTokAppName:              res.Metrics.TikTokAppName,
				//MobileAppID:                res.Metrics.MobileAppID,
				//ImageMode:                  res.Metrics.ImageMode,
				//OnsiteDownloadStart:        res.Metrics.OnsiteDownloadStart,
				//CostPerOnsiteDownloadStart: res.Metrics.CostPerOnsiteDownloadStart,
				//OnsiteDownloadStartRate:    res.Metrics.OnsiteDownloadStartRate,
			}
			extractedModels = append(extractedModels, report)
			return extractedModels, nil
		},
	})
	if err != nil {
		return err
	}

	return extractor.Execute()
}

var ExtractReportMeta = plugin.SubTaskMeta{
	Name:             "ExtractReport",
	EntryPoint:       ExtractReport,
	EnabledByDefault: true,
	Description:      "Extract raw data into tool layer table tiktokads_report",
	DomainTypes:      []string{},
}
