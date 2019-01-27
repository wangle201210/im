package models

import (
	"github.com/astaxie/beego/orm"
	"im/controllers"
	"time"
)

type Pic struct {
	Id       	int64    	`json:"id" orm:"column(id);auto;size(100)"`
	CId      	int64    	`json:"c_id" orm:"column(c_id);size(100);"`
	Url      	string  	`json:"url" orm:"column(url);size(100)"`
	Order    	int64   	`json:"order" orm:"column(order);size(100)"`
	Room     	int64    	`json:"room" orm:"column(room);size(100)"`
	CreatedAt   time.Time	`json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time	`json:"updated_at" orm:"auto_now;type(datetime)"`
}

// User database CRUD methods include Insert, Read, Update and Delete
func (reg *Pic) Insert() error {
	if _, err := orm.NewOrm().Insert(reg); err != nil {
		return err
	}
	controllers.BroadcastPic2All()
	return nil
}

func (reg *Pic) Read(fields ...string) error {
	if err := orm.NewOrm().Read(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Pic) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(reg, fields...); err != nil {
		return err
	}
	controllers.BroadcastPic2All()
	return nil
}

func (reg *Pic) Delete() error {
	if _, err := orm.NewOrm().Delete(reg); err != nil {
		return err
	}
	controllers.BroadcastPic2All()
	return nil
}


func init() {
	orm.RegisterModel(new(Pic))
}