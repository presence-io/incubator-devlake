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
	"github.com/apache/incubator-devlake/core/errors"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

type TiktokAdsApiParams struct {
	ConnectionId uint64
	AdvertiserId string
}

type TiktokAdsOptions struct {
	// TODO add some custom options here if necessary
	// options means some custom params required by plugin running.
	// Such As How many rows do your want
	// You can use it in subtasks, and you need to pass it to main.go and pipelines.
	ConnectionId  uint64                   `json:"connectionId"`
	Name          string                   `json:"name"`
	Tasks         []string                 `json:"tasks,omitempty"`
	AdvertiserID  string                   `json:"advertiserId"`
	ServiceType   string                   `json:"service_type,omitempty"`
	ReportType    string                   `json:"reportType"`
	DataLevel     string                   `json:"data_level,omitempty"`
	Dimensions    []string                 `json:"dimensions"`
	Metrics       []string                 `json:"metrics,omitempty"`
	StartDate     string                   `json:"start_date,omitempty"`
	EndDate       string                   `json:"end_date,omitempty"`
	QueryLifetime bool                     `json:"query_lifetime,omitempty"`
	OrderField    string                   `json:"order_field,omitempty"`
	OrderType     string                   `json:"order_type,omitempty"`
	Filtering     []map[string]interface{} `json:"filtering,omitempty"`
	QueryMode     string                   `json:"query_mode,omitempty"`

	RuleLevel string `json:"rule_level,omitempty"`
}

type TiktokAdsTaskData struct {
	Options   *TiktokAdsOptions
	ApiClient *helper.ApiAsyncClient
}

func DecodeAndValidateTaskOptions(options map[string]interface{}) (*TiktokAdsOptions, errors.Error) {
	var op TiktokAdsOptions
	if err := helper.Decode(options, &op, nil); err != nil {
		return nil, err
	}
	if op.ConnectionId == 0 {
		return nil, errors.Default.New("connectionId is invalid")
	}
	return &op, nil
}
