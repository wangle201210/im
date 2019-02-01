package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/websocket"
	"im/models"
	"strconv"
	"time"
)

type VideoController struct {
	AppController
}

var modVideo models.Video
var modVideoList []models.Video

func (this *VideoController) All() {
	o := orm.NewOrm()
	qs := o.QueryTable("Video")
	all, e := qs.OrderBy("-order").All(&modVideoList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		resp = Response{readSuccess.code,readSuccess.text,modVideoList}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *VideoController) Add() {
	var Video models.Video
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &Video); err == nil {
		err := Video.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,Video}
			BroadcastVideo2All(Video.Room,Video.Url)
		} else {
			resp = Response{addError.code,addError.text,""}
		}
	} else {
		resp = Response{addError.code,addError.text,""}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *VideoController) Delete() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{paraError.code,paraError.text,""}
	} else {
		modVideo.CId = intId
		e := modVideo.Delete("c_id")
		if e != nil {
			resp = Response{deleteError.code,deleteError.text,""}
		} else {
			resp = Response{deleteSuccess.code,deleteSuccess.text,""}
			BroadcastVideo2All(modVideo.Room,modVideo.Url)
		}

	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *VideoController) Edit() {
	var Video models.Video
	var err error
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{updateError.code, updateError.text, ""}
	}
	Video.CId = intId
	read := Video.Read("CId")
	if read != nil {
		resp = Response{updateError.code, updateError.text, ""}
	}
	var getVideo models.Video
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &getVideo)
	if err == nil {
		Video.Url = getVideo.Url
		Video.Room = getVideo.Room
		Video.Content = getVideo.Content
		err := Video.Update()
		if err == nil {
			resp = Response{updateSuccess.code,updateSuccess.text,Video}
			BroadcastVideo2All(Video.Room,Video.Url)
		} else {
			resp = Response{updateError.code,updateError.text,""}
		}
	} else {
		resp = Response{updateError.code,updateError.text,""}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

func (this *VideoController) Show() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{readError.code, readError.text, ""}
	}
	modVideo.CId = intId
	read := modVideo.Read("c_id")
	beego.Info(modVideo)
	if read != nil {
		resp = Response{readError.code, readError.text, ""}
	} else {
		resp = Response{readSuccess.code, readSuccess.text, modVideo}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

// 对应房间视频推给某个人
func broadcastVideo(ws *websocket.Conn, room int64) {
	modVideo.Room = room
	read := modVideo.Read("Room")
	if read != nil {
		beego.Error("video not find")
	}
	beego.Info("video is: ",modVideo)
	event := models.Event{models.EVENT_VIDEO, "system", 0, int(time.Now().UnixNano() / 1e6), modVideo.Url}
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}
	if ws.WriteMessage(websocket.TextMessage, data) != nil {
		unsubscribe <- ws
	}
}

// 所有照片推给所有人
func BroadcastVideo2All(room int64, url string) {
	event := models.Event{models.EVENT_VIDEO, "system", 0, int(time.Now().UnixNano() / 1e6), url}
	broadcastWebSocket(event, room)
}