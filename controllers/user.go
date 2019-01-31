package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"im/helper"
	"im/models"
)

type UserController struct {
	AppController
}

type LoginToken struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

func (this *UserController) Login() {
	req := this.Ctx.Input.RequestBody
	r := models.User{}
	err := json.Unmarshal(req,&r)
	if err != nil {
		beego.Info("json.Unmarshal is err:", err.Error())
	}
	user, b := models.CheckUserAuth(r.Name, r.Password)
	data := Response{}
	if !b {
		//this.Redirect("/", 302)
		data = Response{302,"登陆失败！",""}

	} else {
		et := helper.EasyToken{
			Uid:      user.Id,
			Username: user.Name,
			Password: user.Password,
			Role:     user.Role,
		}
		token, err := et.GetToken()
		if token == "" || err != nil {
			data = Response{302,"登陆失败！",""}
		} else {
			data = Response{200,"登陆成功！", LoginToken{user, token}}
		}
		for _,room := range subscribers{
			if IsUserExist(room, user.Name) {
				data = Response{302,"你的账号已经在别处登陆！",""}
			}
		}
	}
	this.Data["json"] = data
	this.ServeJSON()
	return
}

func (this *UserController) Logout() {
	info, b, _ := this.GetTokenInfo()
	if b {
		ji := info.(map[string]interface{})
		name := helper.Interface2string(ji["Username"])
		uid := helper.Interface2int64(ji["Uid"])
		beego.Info("info is: ",name,uid,ji)
		LogOutLeave(name,uid)//关闭ws
	}
	data := Response{200,"登出成功！",""}
	this.Data["json"] = data
	this.ServeJSON()
	return
}

func (this *UserController) Mine() {
	beego.Info(this.GetTokenInfo())
	this.Data["json"] = ""
	this.ServeJSON()
	return
}

func (this *UserController) GetUserInfo () {
	data := Response{}
	info, b, sub := this.GetTokenInfo()
	if !b {
		data = Response{302,"身份验证失败！",sub}
	} else {
		ji := info.(map[string]interface{})
		id := helper.Interface2int64(ji["Uid"])
		beego.Info(id)
		//id, _ := strconv.ParseInt(i, 10, 64)
		user := models.User{Id:id}
		if read := user.Read("Id");read != nil {
			data = Response{302,"身份验证失败！",""}
		} else {
			data = Response{200,"success", user}
		}
	}
	this.Data["json"] = data
	this.ServeJSON()
	return
}
