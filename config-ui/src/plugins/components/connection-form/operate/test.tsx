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

import { useState, useMemo } from 'react';
import { Button } from '@blueprintjs/core';
import { pick } from 'lodash';

import { operator } from '@/utils';

import * as API from '../api';

interface Props {
  plugin: string;
  values: any;
  errors: any;
}

export const Test = ({ plugin, values, errors }: Props) => {
  const [testing, setTesting] = useState(false);

  const disabled = useMemo(() => {
    return Object.values(errors).some((value) => value);
  }, [errors]);

  const handleSubmit = async () => {
    await operator(
      () =>
        API.testConnection(
          plugin,
          pick(values, [
            'endpoint',
            'token',
            'username',
            'password',
            'proxy',
            'authMethod',
            'appId',
            'secretKey',
            'secret',
            'authCode',
            'tenantId',
            'tenantType',
          ]),
        ),
      {
        setOperating: setTesting,
      },
    );
  };

  return <Button loading={testing} disabled={disabled} outlined text="Test Connection" onClick={handleSubmit} />;
};
