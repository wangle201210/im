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
	qs := o.QueryTable("record")
	all, e := qs.OrderBy("-id").All(&modRecordList)
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
	var record models.Record
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &record); err == nil {
		err := record.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,record}
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
		modRecord.Id = intId
		e := modRecord.Delete()
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
	var record models.Record
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &record); err == nil {
		err := record.Update()
		if err == nil {
			resp = Response{updateSuccess.code,updateSuccess.text,record}
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

func (this *RecordController) Show() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{readError.code, readError.text, ""}
	}
	modRecord.Id = intId
	read := modRecord.Read()
	if read != nil {
		resp = Response{readError.code, readError.text, ""}
	} else {
		resp = Response{readSuccess.code, readSuccess.text, modRecord}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

