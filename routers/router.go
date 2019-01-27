package routers

import (
	"github.com/astaxie/beego"
	"im/controllers"
)

func init() {
	// Register routers.
	beego.Router("/", &controllers.AppController{},"get:Index")
	// Indicate AppController.Join method to handle POST requests.
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/logout", &controllers.UserController{}, "post:Logout")
	beego.Router("/mine", &controllers.UserController{}, "get:Mine")
	beego.Router("/auth/user", &controllers.UserController{}, "get:GetUserInfo")

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
	//pic
	beego.Router("/pics",&controllers.PicController{},"get:All")
	beego.Router("/pics",&controllers.PicController{},"post:Add")
	beego.Router("/pic/:id",&controllers.PicController{},"delete:Delete")
	beego.Router("/pic/:id",&controllers.PicController{},"put:Edit")
	beego.Router("/pic/:id",&controllers.PicController{},"get:Show")
	//record
	beego.Router("/records",&controllers.RecordController{},"get:All")
	beego.Router("/records",&controllers.RecordController{},"post:Add")
	beego.Router("/record/:id",&controllers.RecordController{},"delete:Delete")
	beego.Router("/record/:id",&controllers.RecordController{},"put:Edit")
	beego.Router("/record/:id",&controllers.RecordController{},"get:Show")
	//history
	beego.Router("/histories",&controllers.HistoryController{},"get:All")
	beego.Router("/histories",&controllers.HistoryController{},"post:Add")
	beego.Router("/history/:id",&controllers.HistoryController{},"delete:Delete")
	beego.Router("/history/:id",&controllers.HistoryController{},"put:Edit")
	beego.Router("/history/:id",&controllers.HistoryController{},"get:Show")
	//chat
	beego.Router("/chats",&controllers.ChatController{},"get:All")
	beego.Router("/chats",&controllers.ChatController{},"post:Add")
	beego.Router("/chat/:id",&controllers.ChatController{},"delete:Delete")
	beego.Router("/chat/:id",&controllers.ChatController{},"put:Edit")
	beego.Router("/chat/:id",&controllers.ChatController{},"get:Show")
	//course
	beego.Router("/courses",&controllers.CourseController{},"get:All")
	beego.Router("/courses",&controllers.CourseController{},"post:Add")
	beego.Router("/course/:id",&controllers.CourseController{},"delete:Delete")
	beego.Router("/course/:id",&controllers.CourseController{},"put:Edit")
	beego.Router("/course/:id",&controllers.CourseController{},"get:Show")

}
