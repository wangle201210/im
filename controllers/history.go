package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"im/models"
	"strconv"
)

type HistoryController struct {
	AppController
}

var modHistory models.History
var modHistoryList []models.History

func (this *HistoryController) All() {
	o := orm.NewOrm()
	qs := o.QueryTable("History")
	by := qs.OrderBy("-id")
	s := this.GetString("room")
	if s != "" {
		by = by.Filter("room", s)
	}
	e := by.One(&modHistoryList)
	//all, e := by.All(&modHistoryList)
	if e != nil {
		beego.Info(e)
	} else {
		//beego.Info(all)
		resp = Response{readSuccess.code,readSuccess.text,modHistoryList}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *HistoryController) Add() {
	var History models.History
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &History); err == nil {
		err := History.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,History}
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


func (this *HistoryController) Delete() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{paraError.code,paraError.text,""}
	} else {
		modHistory.CId = intId
		e := modHistory.Delete("c_id")
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


func (this *HistoryController) Edit() {
	var History models.History
	var err error
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{updateError.code, updateError.text, ""}
	}
	History.CId = intId
	read := History.Read("CId")
	if read != nil {
		resp = Response{updateError.code, updateError.text, ""}
	} else {
		var getHistory models.History
		err = json.Unmarshal(this.Ctx.Input.RequestBody, &getHistory)
		if err == nil {
			History.Title = getHistory.Title
			History.Room = getHistory.Room
			History.Content = getHistory.Content
			err := History.Update()
			if err == nil {
				resp = Response{updateSuccess.code,updateSuccess.text,History}
			} else {
				resp = Response{updateError.code,updateError.text,""}
			}
		} else {
			resp = Response{updateError.code,updateError.text,""}
		}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

func (this *HistoryController) Show() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{readError.code, readError.text, ""}
	}
	modHistory.CId = intId
	read := modHistory.Read("c_id")
	beego.Info(modHistory)
	if read != nil {
		resp = Response{readError.code, readError.text, ""}
	} else {
		resp = Response{readSuccess.code, readSuccess.text, modHistory}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

