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

import styled from 'styled-components';

export const Wrapper = styled.div`
  height: 100vh;
  background-color: #f5f5f5;
  overflow: hidden;
`;

export const Inner = styled.div`
  margin: 200px auto;
  padding: 16px 24px;
  width: 400px;
  background-color: #fff;
  border-raidus: 4px;
  box-shadow: 0px 2.4px 4.8px -0.8px rgba(0, 0, 0, 0.1), 0px 1.6px 8px rgba(0, 0, 0, 0.07);

  h2 {
    margin-bottom: 16px;
    text-align: center;
  }
`;
