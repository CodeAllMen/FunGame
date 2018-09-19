package routers

import (
	"github.com/MobileCPX/FunGame/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/?:sp/?:country/?:type", &controllers.PageController{})
	beego.Router("/game/?:id", &controllers.GamePage{})

	beego.Router("/user/?:method", &controllers.UserController{})
}
