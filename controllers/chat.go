package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"im/models"
	"strconv"
)

type ChatController struct {
	AppController
}

var modChat models.Chat
var modChatList []models.Chat

func (this *ChatController) All() {
	o := orm.NewOrm()
	qs := o.QueryTable("chat")
	by := qs.OrderBy("-id")
	s := this.GetString("room")
	if s != "" {
		by = by.Filter("room", s)
	}
	all, e := by.All(&modChatList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		resp = Response{readSuccess.code,readSuccess.text,modChatList}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *ChatController) Add() {
	var chat models.Chat
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &chat); err == nil {
		err := chat.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,chat}
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


func (this *ChatController) Delete() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{paraError.code,paraError.text,""}
	} else {
		modChat.Id = intId
		e := modChat.Delete()
		if e != nil {
			resp = Response{deleteError.code,deleteError.text,""}
		} else {
			resp = Response{deleteSuccess.code,deleteSuccess.text,""}
		}

	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *ChatController) Edit() {
	var chat models.Chat
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &chat); err == nil {
		err := chat.Update()
		if err == nil {
			resp = Response{updateSuccess.code,updateSuccess.text,chat}
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

func (this *ChatController) Show() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{readError.code, readError.text, ""}
	}
	modChat.Id = intId
	read := modChat.Read()
	if read != nil {
		resp = Response{readError.code, readError.text, ""}
	} else {
		resp = Response{readSuccess.code, readSuccess.text, modChat}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

func InsertChat(info *models.Chat) error {
	if e := info.Insert(); e != nil {
		return e
	} else {
		return nil
	}
}
