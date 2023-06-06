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

import { toast } from '@/components';
import { operator } from '@/utils';

import * as API from './api';

interface Props {
  name: string;
  enableDora: boolean;
  onHideDialog: () => void;
}

export const useProject = <T>({ name, enableDora, onHideDialog }: Props) => {
  const [loading, setLoading] = useState(false);
  const [operating, setOperating] = useState(false);
  const [projects, setProjects] = useState<T[]>([]);

  const getProjects = async () => {
    setLoading(true);
    try {
      const res = await API.getProjects({ page: 1, pageSize: 200 });
      setProjects(
        res.projects.map((it: any) => ({
          name: it.name,
        })),
      );
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getProjects();
  }, []);

  const handleSave = async () => {
    if (!/^(\w|-|\/)+$/.test(name)) {
      toast.error('Please enter alphanumeric or underscore');
      return;
    }

    const payload = {
      name,
      description: '',
      metrics: [
        {
          pluginName: 'dora',
          pluginOption: '',
          enable: enableDora,
        },
      ],
    };

    const [success] = await operator(() => API.createProject(payload), {
      setOperating,
      formatReason: () => 'This project name is been used. Please use a new one.',
    });

    if (success) {
      onHideDialog();
      getProjects();
    }
  };

  return useMemo(
    () => ({
      loading,
      operating,
      projects,
      onSave: handleSave,
    }),
    [loading, operating, projects, name, enableDora],
  );
};
