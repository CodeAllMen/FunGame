package dimoco_pl

import (
	"fmt"

	"github.com/astaxie/beego"
)

type SubResultControllers struct {
	beego.Controller
}

func (c *SubResultControllers) Get() {
	// subStatus := c.GetString("status")
	// subID := c.GetString("subID")
	parm := c.Ctx.Request.URL.String()
	fmt.Println(parm)
	subStatus := c.GetString("sph-x")
	subID := c.GetString("sph-s")

	if subStatus == "s" {
		c.SetSession("user", subID)
		c.Data["text"] = "Sukces subskrypcji"
		c.Data["URL"] = "/"
	} else {
		c.Data["text"] = "Nieudana subskrypcja"
		c.Data["URL"] = "/dm/pl/lp"
	}
	c.TplName = "dm/pl/sub_result.html"
}
