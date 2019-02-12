package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"im/models"
	"strconv"
)

type RecordController struct {
	AppController
}

var modRecord models.Record
var modRecordList []models.Record

func (this *RecordController) All() {
	o := orm.NewOrm()
	qs := o.QueryTable("Record")
	by := qs.OrderBy("-id")
	s := this.GetString("room")
	if s != "" {
		by = by.Filter("room", s)
	}
	all, e := by.All(&modRecordList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		resp = Response{readSuccess.code,readSuccess.text,modRecordList}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *RecordController) Add() {
	var Record models.Record
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &Record); err == nil {
		err := Record.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,Record}
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


func (this *RecordController) Delete() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{paraError.code,paraError.text,""}
	} else {
		modRecord.CId = intId
		e := modRecord.Delete("c_id")
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


func (this *RecordController) Edit() {
	var Record models.Record
	var err error
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{updateError.code, updateError.text, ""}
	}
	Record.CId = intId
	read := Record.Read("CId")
	if read != nil {
		resp = Response{updateError.code, updateError.text, ""}
	} else {
		var getRecord models.Record
		err = json.Unmarshal(this.Ctx.Input.RequestBody, &getRecord)
		if err == nil {
			Record.Title = getRecord.Title
			Record.Room = getRecord.Room
			Record.Content = getRecord.Content
			err := Record.Update()
			if err == nil {
				resp = Response{updateSuccess.code,updateSuccess.text,Record}
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

func (this *RecordController) Show() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{readError.code, readError.text, ""}
	}
	modRecord.CId = intId
	read := modRecord.Read("c_id")
	beego.Info(modRecord)
	if read != nil {
		resp = Response{readError.code, readError.text, ""}
	} else {
		resp = Response{readSuccess.code, readSuccess.text, modRecord}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

