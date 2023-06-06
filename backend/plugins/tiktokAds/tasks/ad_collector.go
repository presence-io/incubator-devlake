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
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"net/url"
	"strconv"
	"time"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
)

const RAW_AD_TABLE = "tiktokad_ad"

var _ plugin.SubTaskEntryPoint = CollectAd

func CollectAd(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	collector, err := api.NewApiCollector(api.ApiCollectorArgs{
		RawDataSubTaskArgs: api.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: TiktokAdsApiParams{
				ConnectionId: data.Options.ConnectionId,
				AdvertiserId: data.Options.AdvertiserID,
				StatTimeDate: time.Now().Format("2006-01-02 00:00:00"),
			},
			Table: RAW_AD_TABLE,
		},
		PageSize:      100,
		ApiClient:     data.ApiClient,
		Incremental:   false,
		UrlTemplate:   "v1.3/ad/get/",
		GetTotalPages: GetTotalPages,
		Query: func(reqData *api.RequestData) (url.Values, errors.Error) {

			query := url.Values{}
			query.Set("advertiser_id", data.Options.AdvertiserID)
			// Filtering
			// Set the filtering field as a query parameter
			for _, filter := range data.Options.Filtering {
				filterJSON, _ := json.Marshal(filter)
				query.Add("filtering", string(filterJSON))
			}
			//QueryMode
			query.Set("page", strconv.Itoa(reqData.Pager.Page))
			query.Set("page_size", strconv.Itoa(reqData.Pager.Size))
			return query, nil
		},
		ResponseParser: GetRawMessageArrayFromResponse,
	})

	if err != nil {
		return err
	}
	return collector.Execute()
}

var CollectAdMeta = plugin.SubTaskMeta{
	Name:             "CollectAd",
	EntryPoint:       CollectAd,
	EnabledByDefault: true,
	Description:      "Collect Ad data from Tiktokad api",
	DomainTypes:      []string{},
}
