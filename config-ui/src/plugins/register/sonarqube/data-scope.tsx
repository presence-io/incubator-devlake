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

import React, { useMemo } from 'react';

import { DataScopeMillerColumns } from '@/plugins';

import type { SonarQubeScopeType } from './types';

interface Props {
  connectionId: ID;
  selectedItems: SonarQubeScopeType[];
  onChangeItems: (selectedItems: SonarQubeScopeType[]) => void;
}

export const SonarQubeDataScope = ({ connectionId, onChangeItems, ...props }: Props) => {
  const selectedItems = useMemo(
    () => props.selectedItems.map((it) => ({ id: it.projectKey, data: it })),
    [props.selectedItems],
  );

  return (
    <>
      <h4>Add Repositories by Selecting from the Directory</h4>
      <p>The following directory lists out all projects from SonarQube.</p>
      <DataScopeMillerColumns
        columnCount={1}
        title="Projects"
        plugin="sonarqube"
        connectionId={connectionId}
        selectedItems={selectedItems}
        onChangeItems={onChangeItems}
      />
    </>
  );
};
