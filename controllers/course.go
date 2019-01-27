package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"im/models"
	"strconv"
)

type CourseController struct {
	AppController
}

var modCourse models.Course
var modCourseList []models.Course

func (this *CourseController) All() {
	o := orm.NewOrm()
	qs := o.QueryTable("course")
	all, e := qs.OrderBy("-id").All(&modCourseList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		resp = Response{readSuccess.code,readSuccess.text,modCourseList}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *CourseController) Add() {
	var course models.Course
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &course); err == nil {
		err := course.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,course}
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


func (this *CourseController) Delete() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{paraError.code,paraError.text,""}
	} else {
		modCourse.Id = intId
		e := modCourse.Delete()
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


func (this *CourseController) Edit() {
	var course models.Course
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &course); err == nil {
		err := course.Update()
		if err == nil {
			resp = Response{updateSuccess.code,updateSuccess.text,course}
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

func (this *CourseController) Show() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{readError.code, readError.text, ""}
	}
	modCourse.Id = intId
	read := modCourse.Read()
	if read != nil {
		resp = Response{readError.code, readError.text, ""}
	} else {
		resp = Response{readSuccess.code, readSuccess.text, modCourse}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

