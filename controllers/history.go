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

var modHis models.History
var modHisList []models.History

func (this *HistoryController) All() {
	o := orm.NewOrm()
	qs := o.QueryTable("history")
	all, e := qs.OrderBy("-id").All(&modHisList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		resp = Response{readSuccess.code,readSuccess.text,modHisList}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *HistoryController) Add() {
	var history models.History
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &history); err == nil {
		err := history.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,history}
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
		modHis.Id = intId
		e := modHis.Delete()
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
	var history models.History
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &history); err == nil {
		err := history.Update()
		if err == nil {
			resp = Response{updateSuccess.code,updateSuccess.text,history}
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

func (this *HistoryController) Show() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{readError.code, readError.text, ""}
	}
	modHis.Id = intId
	read := modHis.Read()
	if read != nil {
		resp = Response{readError.code, readError.text, ""}
	} else {
		resp = Response{readSuccess.code, readSuccess.text, modHis}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

