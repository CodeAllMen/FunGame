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

	//页面
	beego.Router("/l", &controllers.LpPage{})
	beego.Router("/h", &controllers.HelpPage{})
	beego.Router("/p", &controllers.PrivacyPage{})
	beego.Router("/t", &controllers.TNCPage{})
}
