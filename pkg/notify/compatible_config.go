// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package notify

import (
	"context"
	"net/http"

	"yunion.io/x/jsonutils"

	"yunion.io/x/onecloud/pkg/appsrv"
	"yunion.io/x/onecloud/pkg/httperrors"
)

const (
	EMAIL_KEYWORDPLURAL = "email_configs"
	EMAIL_KEYWORD       = "email_config"
	EMAIL               = "email"

	SMS_KEYWORDPLURAL = "sms_configs"
	SMS_KEYWORD       = "sms_config"
	SMS               = "mobile"
)

func emailConfigDeleteHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// do not need modify
	manager, _, _, _ := fetchEnv(ctx, w, r)
	params := map[string]string{
		"<type>": EMAIL,
	}
	err := manager.DeleteConfig(ctx, params)
	if err != nil {
		httperrors.GeneralServerError(w, err)
	}
}

func emailConfigGetHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	manager, _, query, _ := fetchEnv(ctx, w, r)
	params := map[string]string{
		"<type>": EMAIL,
	}
	ret, err := manager.GetConfig(ctx, params, query)
	if err != nil {
		httperrors.GeneralServerError(w, err)
	}
	data, _ := ret.Get("config")
	dataDict := data.(*jsonutils.JSONDict)
	output := jsonutils.NewDict()
	output.Add(dataDict, EMAIL_KEYWORD)
	appsrv.SendJSON(w, output)
}

func emailConfigUpdateHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	manager, _, _, body := fetchEnv(ctx, w, r)
	body, _ = body.Get(EMAIL_KEYWORD)
	bodyRet := jsonutils.DeepCopy(body)
	newBody := jsonutils.NewDict()
	newBody.Add(body, EMAIL)
	err := manager.UpdateConfig(ctx, newBody)
	if err != nil {
		httperrors.GeneralServerError(w, err)
	}
	data := jsonutils.NewDict()
	data.Add(jsonutils.NewInt(200), "code")
	data.Add(jsonutils.NewString("OK"), "title")
	data.Add(bodyRet, "message")
	ret := jsonutils.NewDict()
	ret.Add(data, EMAIL_KEYWORD)
	appsrv.SendJSON(w, ret)
}

func smsConfigDeleteHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// do not need modify
	manager, _, _, _ := fetchEnv(ctx, w, r)
	params := map[string]string{
		"<type>": SMS,
	}
	err := manager.DeleteConfig(ctx, params)
	if err != nil {
		httperrors.GeneralServerError(w, err)
	}
}

func smsConfigGetHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	manager, _, query, _ := fetchEnv(ctx, w, r)
	params := map[string]string{
		"<type>": SMS,
	}
	ret, err := manager.GetConfig(ctx, params, query)
	if err != nil {
		httperrors.GeneralServerError(w, err)
	}
	// modify
	data, _ := ret.Get("config")
	appsrv.SendJSON(w, jsonutils.Marshal(map[string]jsonutils.JSONObject{
		SMS_KEYWORD: data,
	}))
}

func smsConfigUpdateHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	manager, _, _, body := fetchEnv(ctx, w, r)
	body, _ = body.Get(SMS_KEYWORD)
	newBody := jsonutils.NewDict()
	newBody.Add(body, SMS)
	err := manager.UpdateConfig(ctx, newBody)
	if err != nil {
		httperrors.GeneralServerError(w, err)
	}
	data := jsonutils.NewDict()
	data.Add(jsonutils.NewInt(200), "code")
	data.Add(jsonutils.NewString("OK"), "title")
	data.Add(body, "message")
	ret := jsonutils.NewDict()
	ret.Add(data, SMS_KEYWORD)
	appsrv.SendJSON(w, ret)
}
