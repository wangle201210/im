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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"im/models"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	baseController
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Get() {
	// Safe check.
	//uname := this.GetString("uname")
	//if len(uname) == 0 {
	//	this.Redirect("/", 302)
	//	return
	//}
	this.TplName = "im.html"
	this.Data["IsWebSocket"] = true
	//this.Data["UserName"] = uname
}
// 接收发送过来的消息（长链接）
// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) Join() {
	uname := this.GetString("uname")
	password := this.GetString("password")
	room, e := this.GetInt64("room")
	if e != nil || room == 0{
		room = 1
	}
	user := models.User{0,uname,password,""}
	u := models.Find(user)
	fmt.Println(u)
	if u.Name == "" {//查看数据库内是否有这个人,
		this.Redirect("/", 302)
		return
	}
	if u.Name != "admin" && !IsUserExist(subscribers[room], "admin") { //管理员不在线
		this.Redirect("/", 302)
		return
	}
	//if IsUserExist(subscribers, uname) {
	//	publish <- newEvent(models.EVENT_OLD, uname, "")
	//}
	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	// 来的时候加个消息推送，断开链接加个离开推送
	// Join chat room.
	Join(uname,room, ws)
	defer Leave(ws)

	// 一直监听前端来的消息
	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		publish <- newEvent(models.EVENT_MESSAGE, uname,room, string(p))
	}
}

// 广播给全部人
// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(event models.Event,room int64) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}
	beego.Info("room is:",room)
	for sub := subscribers[room].Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				// User disconnected.
				unsubscribe <- sub.Value.(Subscriber).Conn
			}
		}
	}
}
