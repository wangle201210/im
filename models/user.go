package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"im/helper"
	"time"
)

type User struct {
	Id       	int64		`json:"id" orm:"column(id);auto;"`
	Name     	string    	`json:"name" orm:"column(name);size(100)"`
	Password 	string    	`json:"password" orm:"column(password);size(100)"`
	Role     	string    	`json:"role" orm:"column(role);size(100)"`
	Token     	string   	`json:"token" orm:"column(token);size(100);null"`
	Chat        []*Chat 	`orm:"reverse(many)"` // 设置一对多的反向关系
	CreatedAt   time.Time	`json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time	`json:"updated_at" orm:"auto_now;type(datetime)"`
}


func (reg *User) SetToken() error {
	token := helper.RandSeq(10)
	reg.Token = token
	beego.Info(token)
	if err := orm.NewOrm().Read(reg,"Token"); err == nil {
		beego.Info("bkn")
		err := reg.SetToken()
		if err != nil {
			return err
		}
	} else {
		err := reg.Update("token")
		if err != nil {
			return err
		}
	}
	return nil
}

func Find(m User) (r User) {
	read := m.Read("Name", "Password")
	if read != nil {
		beego.Debug(read)
	}
	r = m
	return
}

func (m User) IsAdmin(s string) bool {
	read := m.Read(s)
	if read != nil {
		beego.Debug(read)
		return false
	}
	if m.Role == "admin" {
		return true
	}
	return false
}
func CheckUserAuth(name string, password string) (User, bool) {
	o := orm.NewOrm()
	user := User{
		Name: name,
		Password: password,
	}
	err := o.Read(&user, "Name", "Password")
	if err != nil {
		return user, false
	}
	return user, true
}

// User database CRUD methods include Insert, Read, Update and Delete
func (reg *User) Insert() error {
	if _, err := orm.NewOrm().Insert(reg); err != nil {
		return err
	}
	return nil
}

func (reg *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *User) Delete() error {
	if _, err := orm.NewOrm().Delete(reg); err != nil {
		return err
	}
	return nil
}


func init() {
	orm.RegisterModel(new(User))
}