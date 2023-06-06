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

import React, { useContext } from 'react';

import { PageLoading } from '@/components';

import type { ConnectionItemType } from './types';
import type { UseContextValueProps } from './use-context-value';
import { useContextValue } from './use-context-value';

const ConnectionContext = React.createContext<{
  connections: ConnectionItemType[];
  onRefresh: () => void;
  onTest: (selectedConnection: ConnectionItemType) => void;
}>({
  connections: [],
  onRefresh: () => {},
  onTest: () => {},
});

interface Props extends UseContextValueProps {
  children?: React.ReactNode;
}

export const ConnectionContextProvider = ({ children, ...props }: Props) => {
  const { loading, connections, onRefresh, onTest } = useContextValue({
    ...props,
  });

  if (loading) {
    return <PageLoading />;
  }

  return (
    <ConnectionContext.Provider
      value={{
        connections,
        onRefresh,
        onTest,
      }}
    >
      {children}
    </ConnectionContext.Provider>
  );
};

export const ConnectionContextConsumer = ConnectionContext.Consumer;

export const useConnection = () => useContext(ConnectionContext);
