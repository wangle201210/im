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
	qs := o.QueryTable("Course")
	by := qs.OrderBy("-id")
	s := this.GetString("room")
	if s != "" {
		by = by.Filter("room", s)
	}
	all, e := by.All(&modCourseList)
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
	var Course models.Course
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &Course); err == nil {
		err := Course.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,Course}
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
		modCourse.CId = intId
		e := modCourse.Delete("c_id")
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
	var Course models.Course
	var err error
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{updateError.code, updateError.text, ""}
	}
	Course.CId = intId
	read := Course.Read("CId")
	if read != nil {
		resp = Response{updateError.code, updateError.text, ""}
	} else {
		var getCourse models.Course
		err = json.Unmarshal(this.Ctx.Input.RequestBody, &getCourse)
		if err == nil {
			Course.Room = getCourse.Room
			Course.Content = getCourse.Content
			err := Course.Update()
			if err == nil {
				resp = Response{updateSuccess.code,updateSuccess.text,Course}
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

func (this *CourseController) Show() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{readError.code, readError.text, ""}
	}
	modCourse.CId = intId
	read := modCourse.Read("c_id")
	beego.Info(modCourse)
	if read != nil {
		resp = Response{readError.code, readError.text, ""}
	} else {
		resp = Response{readSuccess.code, readSuccess.text, modCourse}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}

