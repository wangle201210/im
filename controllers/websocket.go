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
	"im/helper"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"im/models"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	AppController
}

// 接收发送过来的消息（长链接）
// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) Join() {
	token := this.GetString("token")
	room, e := this.GetInt64("room")
	if e != nil || room == 0{
		room = 1
	}
	info, b, sub := this.GetTokenInfo(token)
	beego.Info(info,b,sub)
	ji := info.(map[string]interface{})
	uname := helper.Interface2string (ji["Username"])
	password := helper.Interface2string (ji["Password"])
	user := models.User{}
	user.Name = uname
	user.Password = password
	e = user.Read("Name", "Password")
	if e != nil {//查看数据库内是否有这个人,
		this.Redirect("/", 302)
		return
	}
	if user.Role != "admin" && !IsAdminExist(subscribers[room]) { //管理员不在线
		this.Redirect("/", 302)
		return
	}
	//if user.Role == "admin"  && IsAdminExist(subscribers[room]) {
	//	this.Redirect("/", 302)
	//	return
	//}
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
	broadcastPics(ws)
	broadcastVideo(ws,room)
	// 一直监听前端来的消息
	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		chat := models.Chat{}
		chat.Room = room
		chat.User = &user
		chat.Content = string(p)
		err = InsertChat(&chat)
		if err != nil {
			beego.Info("chat insert err:",err)
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
	//beego.Info("room is:",room)
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
