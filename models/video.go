package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Video struct {
	Id       	int64    	`json:"id" orm:"column(id);auto;size(100)"`
	CId      	int64    	`json:"c_id" orm:"column(c_id);size(100);"`
	Url     	string  	`json:"url" orm:"column(url);type(text)"`
	Content     string  	`json:"content" orm:"column(content);type(text)"`
	Room     	int64    	`json:"room" orm:"column(room);size(100)"`
	CreatedAt   time.Time	`json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time	`json:"updated_at" orm:"auto_now;type(datetime)"`
}

// User database CRUD methods include Insert, Read, Update and Delete
func (reg *Video) Insert() error {
	if _, err := orm.NewOrm().Insert(reg); err != nil {
		return err
	}
	return nil
}

func (reg *Video) Read(fields ...string) error {
	if err := orm.NewOrm().Read(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Video) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(reg, fields...); err != nil {
		return err
	}
	return nil
}

//添加可按fields 条件删除
func (reg *Video) Delete(fields ...string) error {
	read := reg.Read(fields...)
	if read != nil {
		return  read
	}
	if _, err := orm.NewOrm().Delete(reg); err != nil {
		return err
	}
	return nil
}


func init() {
	orm.RegisterModel(new(Video))
}