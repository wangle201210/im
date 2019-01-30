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

type PicController struct {
	AppController
}

var modPic models.Pic
var modPicList []models.Pic

func (this *PicController) All() {
	o := orm.NewOrm()
	qs := o.QueryTable("pic")
	all, e := qs.OrderBy("-order").All(&modPicList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		resp = Response{readSuccess.code,readSuccess.text,modPicList}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *PicController) Add() {
	var pic models.Pic
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &pic); err == nil {
		err := pic.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,pic}
			BroadcastPic2All()
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


func (this *PicController) Delete() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{paraError.code,paraError.text,""}
	} else {
		modPic.CId = intId
		e := modPic.Delete("c_id")
		if e != nil {
			resp = Response{deleteError.code,deleteError.text,""}
		} else {
			resp = Response{deleteSuccess.code,deleteSuccess.text,""}
			BroadcastPic2All()
		}

	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *PicController) Edit() {
	var pic models.Pic
	var err error
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{updateError.code, updateError.text, ""}
	}
	pic.CId = intId
	read := pic.Read("CId")
	if read != nil {
		resp = Response{updateError.code, updateError.text, ""}
	}
	var getPic models.Pic
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &getPic); err == nil {
		pic.Url = getPic.Url
		pic.Room = getPic.Room
		pic.Order = getPic.Order
		err := pic.Update()
		if err == nil {
			resp = Response{updateSuccess.code,updateSuccess.text,pic}
			BroadcastPic2All()
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

func (this *PicController) Show() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{readError.code, readError.text, ""}
	}
	modPic.CId = intId
	read := modPic.Read("c_id")
	beego.Info(modPic)
	if read != nil {
		resp = Response{readError.code, readError.text, ""}
	} else {
		resp = Response{readSuccess.code, readSuccess.text, modPic}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

// 所有照片推给某个人
func broadcastPics(ws *websocket.Conn)  {
	o := orm.NewOrm()
	qs := o.QueryTable("pic")
	all, e := qs.OrderBy("order").All(&modPicList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		// 把所有照片推给用户
		for _,v := range modPicList {
			event := models.Event{models.EVENT_IMG,"system",0,int(time.Now().UnixNano()/1e6),v.Url}
			data, err := json.Marshal(event)
			if err != nil {
				beego.Error("Fail to marshal event:", err)
				return
			}
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				unsubscribe <- ws
			}
		}
	}
}
// 所有照片推给所有人
func BroadcastPic2All()  {
	return
	o := orm.NewOrm()
	qs := o.QueryTable("pic")
	all, e := qs.OrderBy("order").All(&modPicList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		// 把所有照片推给用户
		for k,v := range modPicList {
			var etype models.EventType
			//beego.Info("img key is :",k)
			if k == 0  {
				etype = models.EVENT_NEWIMG
			} else {
				etype = models.EVENT_IMG
			}
			event := models.Event{etype,"system",0,int(time.Now().UnixNano()/1e6),v.Url}
			count, _ := beego.AppConfig.Int64("roomCount")
			for i := int64(0); i <= count; i++ {
				broadcastWebSocket(event,i)
			}
		}
	}
}