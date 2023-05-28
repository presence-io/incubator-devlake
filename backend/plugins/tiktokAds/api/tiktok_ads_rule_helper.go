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

package api

import (
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/log"
	"github.com/apache/incubator-devlake/core/models"
	plugin "github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/go-playground/validator/v10"
	"strconv"
)

// TiktokAdsRuleApiHelper is used to write the CURD of tiktokAds
type TiktokAdsRuleApiHelper struct {
	encKey    string
	log       log.Logger
	db        dal.Dal
	validator *validator.Validate
}

// NewTiktokAdsHelper creates a TiktokAdsHelper for tiktokAds management
func NewTiktokAdsHelper(
	basicRes context.BasicRes,
	vld *validator.Validate,
) *TiktokAdsRuleApiHelper {
	if vld == nil {
		vld = validator.New()
	}
	return &TiktokAdsRuleApiHelper{
		encKey:    basicRes.GetConfig(plugin.EncodeKeyEnvStr),
		log:       basicRes.GetLogger(),
		db:        basicRes.GetDal(),
		validator: vld,
	}
}

// Create a tiktokAds record based on request body
func (c *TiktokAdsRuleApiHelper) Create(tiktokAds interface{}, input *plugin.ApiResourceInput) errors.Error {
	// update fields from request body
	err := c.merge(tiktokAds, input.Body)
	if err != nil {
		return err
	}
	return c.save(tiktokAds, c.db.Create)
}

// Patch (Modify) a tiktokAds record based on request body
func (c *TiktokAdsRuleApiHelper) Patch(tiktokAds interface{}, input *plugin.ApiResourceInput) errors.Error {
	err := c.First(tiktokAds, input.Params)
	if err != nil {
		return err
	}

	err = c.merge(tiktokAds, input.Body)
	if err != nil {
		return err
	}
	return c.save(tiktokAds, c.db.CreateOrUpdate)
}

// First finds tiktokAds from db  by parsing request input and decrypt it
func (c *TiktokAdsRuleApiHelper) First(tiktokAds interface{}, params map[string]string) errors.Error {
	tiktokAdsId := params["tiktokAdsId"]
	if tiktokAdsId == "" {
		return errors.BadInput.New("missing tiktokAdsId")
	}
	id, err := strconv.ParseUint(tiktokAdsId, 10, 64)
	if err != nil || id < 1 {
		return errors.BadInput.New("invalid tiktokAdsId")
	}
	return c.FirstById(tiktokAds, id)
}

// FirstById finds tiktokAds from db by id and decrypt it
func (c *TiktokAdsRuleApiHelper) FirstById(tiktokAds interface{}, id uint64) errors.Error {
	return api.CallDB(c.db.First, tiktokAds, dal.Where("id = ?", id))
}

// List returns all tiktokAds with password/token decrypted
func (c *TiktokAdsRuleApiHelper) List(tiktokAds interface{}) errors.Error {
	return api.CallDB(c.db.All, tiktokAds)
}

func (c *TiktokAdsRuleApiHelper) merge(tiktokAds interface{}, body map[string]interface{}) errors.Error {
	tiktokAds = models.UnwrapObject(tiktokAds)
	return api.Decode(body, tiktokAds, c.validator)
}

func (c *TiktokAdsRuleApiHelper) save(tiktokAds interface{}, method func(entity interface{}, clauses ...dal.Clause) errors.Error) errors.Error {
	err := api.CallDB(method, tiktokAds)
	if err != nil {
		if c.db.IsDuplicationError(err) {
			return errors.BadInput.New("the tiktokAds name already exists")
		}
		return err
	}
	return nil
}

// ConnectionValidator represents the API Connection would validate its fields with customized logic
type ConnectionValidator interface {
	ValidateConnection(connection interface{}, valdator *validator.Validate) errors.Error
}
