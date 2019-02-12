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

// This sample is about using long polling and WebSocket to build a web-based chat room based on beego.
package main

import (
	"github.com/astaxie/beego"
	"im/models"
	_ "im/routers"
)

const (
	APP_VER = "1.0"
)

func main() {
	beego.SetStaticPath("/im/pic", "static/pic")
	beego.SetStaticPath("/videos", "static/videos")
	beego.SetStaticPath("/js", "views/js")
	beego.SetStaticPath("/img", "views/img")
	beego.SetStaticPath("/fonts", "views/fonts")
	beego.SetStaticPath("/css", "views/css")
	beego.Run()
}
func init() {
	models.Init()
}
