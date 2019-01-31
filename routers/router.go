package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"im/controllers"
	"im/helper"
	"strings"
)

func init() {
	// Register routers.
	beego.Router("/", &controllers.AppController{},"get:Index")
	// Indicate AppController.Join method to handle POST requests.
	beego.Router("/api/login", &controllers.UserController{}, "post:Login")
	beego.Router("/api/logout", &controllers.UserController{}, "post:Logout")
	beego.Router("/api/mine", &controllers.UserController{}, "get:Mine")
	beego.Router("/api/auth/user", &controllers.UserController{}, "get:GetUserInfo")

	ns := beego.NewNamespace("/api",
		//beego.NSBefore(Auth),
		beego.NSCond(func(ctx *context.Context) bool {
			return true
		}),
		beego.NSRouter("/pics",&controllers.PicController{},"get:All"),
		beego.NSRouter("/records",&controllers.RecordController{},"get:All"),
		beego.NSRouter("/histories",&controllers.HistoryController{},"get:All"),
		beego.NSRouter("/chats",&controllers.ChatController{},"get:All"),
		beego.NSRouter("/courses",&controllers.CourseController{},"get:All"),
		beego.NSRouter("/videos",&controllers.VideoController{},"get:All"),

		//beego.NSRouter("/ws", &controllers.WebSocketController{}),
		beego.NSRouter("/ws/join", &controllers.WebSocketController{}, "get:Join"),
		beego.NSNamespace("/admin",
			beego.NSBefore(IsAdmin),
			//pic
			beego.NSRouter("/pics",&controllers.PicController{},"post:Add"),
			beego.NSRouter("/pic/:id",&controllers.PicController{},"delete:Delete"),
			beego.NSRouter("/pic/:id",&controllers.PicController{},"put:Edit"),
			beego.NSRouter("/pic/:id",&controllers.PicController{},"get:Show"),
			//record
			beego.NSRouter("/records",&controllers.RecordController{},"post:Add"),
			beego.NSRouter("/record/:id",&controllers.RecordController{},"delete:Delete"),
			beego.NSRouter("/record/:id",&controllers.RecordController{},"put:Edit"),
			beego.NSRouter("/record/:id",&controllers.RecordController{},"get:Show"),
			//history
			beego.NSRouter("/histories",&controllers.HistoryController{},"post:Add"),
			beego.NSRouter("/history/:id",&controllers.HistoryController{},"delete:Delete"),
			beego.NSRouter("/history/:id",&controllers.HistoryController{},"put:Edit"),
			beego.NSRouter("/history/:id",&controllers.HistoryController{},"get:Show"),
			//chat
			beego.NSRouter("/chats",&controllers.ChatController{},"post:Add"),
			beego.NSRouter("/chat/:id",&controllers.ChatController{},"delete:Delete"),
			beego.NSRouter("/chat/:id",&controllers.ChatController{},"put:Edit"),
			beego.NSRouter("/chat/:id",&controllers.ChatController{},"get:Show"),
			//course
			beego.NSRouter("/courses",&controllers.CourseController{},"post:Add"),
			beego.NSRouter("/course/:id",&controllers.CourseController{},"delete:Delete"),
			beego.NSRouter("/course/:id",&controllers.CourseController{},"put:Edit"),
			beego.NSRouter("/course/:id",&controllers.CourseController{},"get:Show"),
			//video
			beego.NSRouter("/videos",&controllers.VideoController{},"post:Add"),
			beego.NSRouter("/video/:id",&controllers.VideoController{},"delete:Delete"),
			beego.NSRouter("/video/:id",&controllers.VideoController{},"put:Edit"),
			beego.NSRouter("/video/:id",&controllers.VideoController{},"get:Show"),
		),
	)
	beego.AddNamespace(ns)
	// WebSocket.

}

var Auth = func(ctx *context.Context){
	authString := ctx.Input.Header("Authorization")
	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		ctx.Output.Body([]byte("token is allowed"))
	}
	tokenString := kv[1]
	if _, b, _ := controllers.TokenInfo(tokenString); !b {
		ctx.Output.Body([]byte("token is allowed"))
	}
}

var IsAdmin = func(ctx *context.Context){
	authString := ctx.Input.Header("Authorization")
	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		ctx.Output.Body([]byte("token is allowed"))
	}
	tokenString := kv[1]
	info, b, _ := controllers.TokenInfo(tokenString)
	if !b {
		ctx.Output.Body([]byte("token is allowed"))
	}
	ji := info.(map[string]interface{})
	role := helper.Interface2string (ji["Role"])
	if role != "admin" {
		ctx.Output.Body([]byte("need admin"))
	}
}