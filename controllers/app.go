// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"im/helper"
	"net/http"
	"strings"
)

type Response struct {
	Code int64         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
var (
	verifyKey  string
	resp = Response{}
)

type status struct {
	code int64
	text string
}
var (
	addSuccess		= status{http.StatusCreated,"添加成功"}
	addError		= status{http.StatusNoContent,"添加失败"}
	readSuccess		= status{http.StatusOK,"查询成功"}
	readError		= status{http.StatusNoContent,"查询失败"}
	deleteSuccess	= status{http.StatusOK,"删除成功"}
	deleteError		= status{http.StatusNoContent,"查询失败"}
	updateSuccess	= status{http.StatusResetContent,"更新成功"}
	updateError		= status{http.StatusNoContent,"查询失败"}
	paraError       = status{http.StatusAccepted,"传入参数不对"}
)

type AppController struct {
	beego.Controller
}

func (this *AppController) Index() {
	this.TplName = "index.html"
}

func (this *AppController) GetTokenInfo(token ...string) (info interface{},b bool, sub string) {
	tokenString := ""
	if len(token) > 0 {
		tokenString = token[0]
	} else {
		authString := this.Ctx.Input.Header("Authorization")
		kv := strings.Split(authString, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
			beego.Error("AuthString invalid:", authString)
		}
		tokenString = kv[1]
	}
	info, b ,sub = helper.ParseToken(tokenString,verifyKey)
	return
}

func init() {
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	verifyKey = appConf.String("jwt::token")
	beego.BConfig.WebConfig.TemplateLeft="<<<"
 	beego.BConfig.WebConfig.TemplateRight=">>>"
 }