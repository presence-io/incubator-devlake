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
export interface Report {
  name: string;
  status: string;
  adGroupId: string;
  deliverySuggestions: string;
  adScheduling: string;
  budget: string;
  bid: string;
  totalCost: string;
  cpc: string;
  clicks: string;
  totalAppStoreClickOnsite: string;
  costPerAppStoreClickOnsite: string;
  ctr: string;
  cpm: string;
  impressions: string;
  conversionsSKAN: string;
  cpaSKAN: string;
  cvr: string;
  conversions: string;
  cpa: string;
  cvrSKAN: string;
  result: string;
  costPerResult: string;
  resultRate: string;
}

export const getReports = (timeRange: string): Report[] => {
  // Send request to server to get reports data for the given time range
  // Return the data as an array of Report objects
  return [
    {
      name: 'HL剪辑',
      status: 'Active',
      adGroupId: 'All',
      deliverySuggestions: '',
      adScheduling: '',
      budget: '656.24 USD',
      bid: '',
      totalCost: '0.20 USD',
      cpc: '4.33 USD',
      clicks: '2.13%',
      totalAppStoreClickOnsite: '',
      costPerAppStoreClickOnsite: '',
      ctr: '0',
      cpm: '151,425',
      impressions: '3,230',
      conversionsSKAN: '0.00 USD',
      cpaSKAN: '0.00%',
      cvr: '21.27%',
      conversions: '687',
      cpa: '0.96 USD',
      cvrSKAN: '',
      result: '',
      costPerResult: '',
      resultRate: '',
    },
    // Add more reports here
  ];
};

export const renderReportsTable = (reports: Report[]): void => {
  const table = document.querySelector('table tbody');
  if (table) {
    table.innerHTML = '';
    reports.forEach((report) => {
      const row = document.createElement('tr');
      row.innerHTML = `
        <td>${report.name}</td>
        <td>${report.status}</td>
        <td>${report.adGroupId}</td>
        <td>${report.deliverySuggestions}</td>
        <td>${report.adScheduling}</td>
        <td>${report.budget}</td>
        <td>${report.bid}</td>
        <td>${report.totalCost}</td>
        <td>${report.cpc}</td>
        <td>${report.clicks}</td>
        <td>${report.totalAppStoreClickOnsite}</td>
        <td>${report.costPerAppStoreClickOnsite}</td>
        <td>${report.ctr}</td>
        <td>${report.cpm}</td>
        <td>${report.impressions}</td>
        <td>${report.conversionsSKAN}</td>
        <td>${report.cpaSKAN}</td>
        <td>${report.cvr}</td>
        <td>${report.conversions}</td>
        <td>${report.cpa}</td>
        <td>${report.cvrSKAN}</td>
        <td>${report.result}</td>
        <td>${report.costPerResult}</td>
        <td>${report.resultRate}</td>
      `;
      table.appendChild(row);
    });
  }
};

const timeRangeEditor = document.querySelector('.vi-date-editor--daterange');
if (timeRangeEditor) {
  timeRangeEditor.addEventListener('change', () => {
    const startDate = (document.querySelector('#start-date') as HTMLInputElement)?.value;
    const endDate = (document.querySelector('#end-date') as HTMLInputElement)?.value;
    const timeRange = `${startDate} - ${endDate}`;
    const reports = getReports(timeRange);
    renderReportsTable(reports);
  });
}

const tabs = document.querySelectorAll<HTMLButtonElement>('.tab');
const tabContents = document.querySelectorAll<HTMLElement>('.tab-content');

tabs.forEach((tab) => {
  tab.addEventListener('click', () => {
    const tabName = tab.dataset.tab;
    tabs.forEach((t) => t.classList.remove('active'));
    tab.classList.add('active');
    tabContents.forEach((tc) => {
      if (tc.id === tabName) {
        tc.classList.add('active');
      } else {
        tc.classList.remove('active');
      }
    });
  });
});
