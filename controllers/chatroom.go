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
	"container/list"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"im/models"
)

type Subscription struct {
	Archive []models.Event      // All the events from the archive.
	New     <-chan models.Event // New events coming in.
}

func newEvent(ep models.EventType, user string,room int64, msg string) models.Event {
	return models.Event{ep, user, room,int(time.Now().UnixNano()/1e6), msg}
}

func Join(user string,room int64, ws *websocket.Conn) {
	subscribe <- Subscriber{user,room, ws}
}

func Leave(ws *websocket.Conn) {
	//beego.Info("ws is: ",ws)
	unsubscribe <- ws
}

func LogOutLeave(name string,uid int64) {
	beego.Info("name room ",name,uid)
	for k,room := range subscribers {
		for sub := room.Front(); sub != nil; sub = sub.Next() {
			if sub.Value.(Subscriber).Name == name {
				subscribers[k].Remove(sub)
				// Clone connection.
				name := sub.Value.(Subscriber).Name
				sub.Value.(Subscriber).Conn.Close()
				beego.Error("WebSocket closed:", name)
				publish <- newEvent(models.EVENT_LEAVE, sub.Value.(Subscriber).Name,k, "") // Publish a LEAVE event.
				break
			}
		}
	}
}

type Subscriber struct {
	Name string
	Room int64
	Conn *websocket.Conn // Only for WebSocket users; otherwise nil.
}

type subList struct {
	List *list.List
	Element *list.Element
}
var (
	// Channel for new join users.
	subscribe = make(chan Subscriber, 100)
	// Channel for exit users.
	unsubscribe = make(chan *websocket.Conn, 100)
	// Send events here to publish them.
	publish = make(chan models.Event, 100) // 所有发过来的消息
	// Long polling waiting list.
	waitingList = list.New()
	//subscribers = map[int64]*list.List{1: list.New(), 2:list.New()} // 储存订阅者
)

var subscribers = func() ( map[int64]*list.List){
	m := map[int64]*list.List{}
	count, _ := beego.AppConfig.Int64("roomCount")
	for i := int64(0); i <= count; i++ {
		m[i] = list.New()
	}
	return m
}()
// This function handles all incoming chan messages.
func chatroom() {
	for {
		select {
		// 如果来了订阅者相关消息（离开或进来）
		case sub := <-subscribe:
			if  subscribers[sub.Room] == nil || !IsUserExist(subscribers[sub.Room], sub.Name) {
				subscribers[sub.Room].PushBack(sub) // Add user to the end of list.
				// Publish a JOIN event.
				publish <- newEvent(models.EVENT_JOIN, sub.Name,sub.Room, "新人到来")
				beego.Info("New user:", sub.Name, ";WebSocket:", sub.Conn != nil,";room:",sub.Room)
			} else {
				subscribers[sub.Room].PushBack(sub) // Add user to the end of list. 允许同一账号多次登陆
				publish <- newEvent(models.EVENT_OLD, sub.Name, sub.Room,"我又来了！")
				beego.Info("Old user:", sub.Name, ";WebSocket:", sub.Conn != nil,";room:",sub.Room)
			}
		// 如果有人发消息
		case event := <-publish:
			// 倒叙遍历待发送的消息
			// Notify waiting list.
			for ch := waitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}

			broadcastWebSocket(event,event.Room)
			models.NewArchive(event)

			//if event.Type == models.EVENT_MESSAGE {
			//	beego.Info("Message from", event.User, ";Content:", event.Content,";room:",event.Room)
			//}
		case unsub := <-unsubscribe:
			for k,room := range subscribers {
				for sub := room.Front(); sub != nil; sub = sub.Next() {
					if sub.Value.(Subscriber).Conn == unsub {
						subscribers[k].Remove(sub)
						// Clone connection.
						name := sub.Value.(Subscriber).Name
						unsub.Close()
						beego.Error("WebSocket closed:", name)
						publish <- newEvent(models.EVENT_LEAVE, sub.Value.(Subscriber).Name,k, "") // Publish a LEAVE event.
						break
					}
				}
			}

		}
	}
}
// 默认都会先执行init
func init() {
	go chatroom()
}

func IsUserExist(room *list.List, user string) bool {
	for sub := room.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}

func IsAdminExist(room *list.List) bool {
	for sub := room.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == "admin" {
			return true
		}
	}
	return false
}
