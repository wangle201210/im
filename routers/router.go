package routers

import (
	"github.com/astaxie/beego"
	"im/controllers"
)

func init() {
	// Register routers.
	beego.Router("/", &controllers.AppController{})
	// Indicate AppController.Join method to handle POST requests.
	beego.Router("/join", &controllers.AppController{}, "post:Join")

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

}
