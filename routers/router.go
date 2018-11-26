package routers

import (
	"github.com/MobileCPX/FunGame/controllers"
	"github.com/MobileCPX/FunGame/controllers/dimoco_pl"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	//Dimoco
	beego.Router("/dm/pl/lp", &dimoco_pl.PLLPPageControllers{})
	beego.Router("/dm/pl/welcome", &dimoco_pl.SubResultControllers{})
	beego.Router("/sub/result/page", &dimoco_pl.SubResultControllers{})

	beego.Router("/?:sp/?:country/?:type", &controllers.PageController{})
	beego.Router("/game/?:id", &controllers.GamePage{})

	beego.Router("/user/?:method", &controllers.UserController{})

	// beego.Router("/pl/start",&dimoco_pl.)

	//页面
	beego.Router("/l", &controllers.LpPage{})
	beego.Router("/h", &controllers.HelpPage{})
	beego.Router("/p", &controllers.PrivacyPage{})
	beego.Router("/t", &controllers.TNCPage{})
}
