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
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/plugins/tiktokAds/models"
	"strconv"
)

var _ plugin.SubTaskEntryPoint = TurnOnAllAdGroups

func TurnOnAllAdGroups(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*TiktokAdsTaskData)
	enableAdGroup := make([]uint64, 0)

	for _, adGroupId := range data.Options.AdGroupIds {
		uintValue, err := strconv.ParseUint(adGroupId, 10, 64)
		if err != nil {
			return errors.Convert(err)
		}
		enableAdGroup = append(enableAdGroup, uintValue)

	}
	prepareUpdate("adgroup", enableAdGroup, taskCtx, models.ENABLE)

	return nil
}

var TurnOnAllAdGroupsMeta = plugin.SubTaskMeta{
	Name:             "TurnOnAllAdGroups",
	EntryPoint:       TurnOnAllAdGroups,
	EnabledByDefault: true,
	Description:      "",
	DomainTypes:      []string{},
}
