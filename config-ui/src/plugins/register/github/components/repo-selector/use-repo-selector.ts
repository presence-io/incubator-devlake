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

import { useEffect, useMemo, useState } from 'react';

import { useProxyPrefix } from '@/hooks';

import type { ScopeItemType } from '../../types';
import * as API from '../../api';

export interface UseRepoSelectorProps {
  connectionId: ID;
}

export const useRepoSelector = ({ connectionId }: UseRepoSelectorProps) => {
  const [loading, setLoading] = useState(false);
  const [items, setItems] = useState<ScopeItemType[]>([]);
  const [search, setSearch] = useState('');

  const prefix = useProxyPrefix({
    plugin: 'github',
    connectionId,
  });

  useEffect(() => {
    if (!search) return;
    setItems([]);
    setLoading(true);

    const timer = setTimeout(async () => {
      try {
        const res = await API.searchRepo(prefix, { q: `${search} fork:true` });
        setItems(
          res.items.map((it: any) => ({
            connectionId,
            githubId: it.id,
            name: `${it.owner.login}/${it.name}`,
            ownerId: it.owner.id,
            language: it.language,
            description: it.description,
            cloneUrl: it.clone_url,
            HTMLUrl: it.html_url,
          })),
        );
      } finally {
        setLoading(false);
      }
    }, 1000);

    return () => clearTimeout(timer);
  }, [prefix, search]);

  return useMemo(
    () => ({
      loading,
      items,
      onSearch(s: string) {
        setSearch(s);
      },
    }),
    [loading, items],
  );
};
