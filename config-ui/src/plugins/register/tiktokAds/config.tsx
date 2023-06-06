/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import React from 'react';

import { ExternalLink } from '@/components';
import type { PluginConfigType } from '@/plugins';
import { PluginType } from '@/plugins';

import Icon from './assets/icon.svg';

export const TIKTOK_ADSConfig: PluginConfigType = {
  entities: [],
  transformation: undefined,
  type: PluginType.Connection,
  plugin: 'tiktokAds',
  name: 'TIKTOK_ADS',
  icon: Icon,
  sort: 9,
  connection: {
    docLink: 'https://devlake.apache.org/docs/Configuration/TiktokAds',
    initialValues: {
      endpoint: 'https://business-api.tiktok.com/open_api/',
    },
    fields: [
      'name',
      {
        key: 'endpoint',
        subLabel: 'You do not need to enter the endpoint URL, because all versions use the same URL.',
        disabled: true,
      },
      {
        key: 'token',
        label: 'Personal Access Token',
      },
      {
        key: 'rateLimitPerHour',
        subLabel:
          'By default, DevLake uses 3,000 requests/hour for data collection for TIKTOK_ADS. But you can adjust the collection speed by setting up your desirable rate limit.',
        learnMore: 'https://devlake.apache.org/docs/Configuration/TiktokAds#fixed-rate-limit-optional',
        externalInfo: 'The maximum rate limit of TIKTOK_ADS is 3,600 requests/hour.',
        defaultValue: 3000,
      },
    ],
  },
};
