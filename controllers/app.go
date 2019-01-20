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
)


type baseController struct {
	beego.Controller
}


type AppController struct {
	baseController
}

func (this *AppController) Get() {
	this.TplName = "index.html"
}

// Join method handles POST requests for AppController.
func (this *AppController) Join() {
	uname := this.GetString("uname")
	password := this.GetString("password")
	room, _ := this.GetInt64("room")

	// Check valid.
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}
	this.Redirect("/ws/join?uname="+uname+"room="+string(room)+"password="+password, 302)
	return
}

func init() {
	beego.BConfig.WebConfig.TemplateLeft="<<<"
 	beego.BConfig.WebConfig.TemplateRight=">>>"
 }