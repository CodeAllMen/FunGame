package routers

import (
	"github.com/MobileCPX/FunGame/controllers"
	"github.com/MobileCPX/FunGame/controllers/dimoco_pl"
	"github.com/MobileCPX/FunGame/controllers/mondia_dz"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/registere/username", &controllers.RegisteredControllers{}) // 电话号码注册用户

	//Dimoco
	beego.Router("/dm/pl/lp", &dimoco_pl.PLLPPageControllers{})
	beego.Router("/dm/pl/welcome", &dimoco_pl.SubResultControllers{})
	beego.Router("/sub/result/page", &dimoco_pl.SubResultControllers{})

	beego.Router("/?:sp/?:country/?:type", &controllers.PageController{})
	beego.Router("/game/?:id", &controllers.GamePage{})

	beego.Router("/user/?:method", &controllers.UserController{})

	// 波兰退订页面
	beego.Router("/unsub/page", &dimoco_pl.UnsubPageControllers{})
	beego.Router("/unsub/result", &dimoco_pl.UnsubResultControllers{})
	beego.Router("/unsub/msisdn", &dimoco_pl.MsisdnUnsubPageControllers{})

	// beego.Router("/pl/start",&dimoco_pl.)

	//页面
	beego.Router("/l", &controllers.LpPage{})
	beego.Router("/h", &controllers.HelpPage{})
	beego.Router("/p", &controllers.PrivacyPage{})
	beego.Router("/t", &controllers.TNCPage{})

	// mondia dz 阿尔及利亚
	// LP页面
	beego.Router("/mm/dz/lp", &mondia_dz.MondiaDZPageController{}, "Get:LpIndex")
	beego.Router("/mm/dz/terms", &mondia_dz.MondiaDZPageController{}, "Get:Terms")

}
