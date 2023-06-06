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

package main

import (
	"github.com/apache/incubator-devlake/core/config"
	"github.com/apache/incubator-devlake/core/plugin"
	_ "github.com/apache/incubator-devlake/core/version"
	"github.com/apache/incubator-devlake/server/api"
)

func main() {
	v := config.GetConfig()
	encKey := v.GetString(plugin.EncodeKeyEnvStr)
	if encKey == "" {
		// Randomly generate a bunch of encryption keys and set them to config
		encKey = plugin.RandomEncKey()
		v.Set(plugin.EncodeKeyEnvStr, encKey)
		err := config.WriteConfig(v)
		if err != nil {
			panic(err)
		}
	}
	api.CreateApiService()
}
