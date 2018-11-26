package dimoco_pl

import (
	"github.com/astaxie/beego"
)

type PLLPPageControllers struct {
	beego.Controller
}

func (c *PLLPPageControllers) Get() {
	trackID := c.GetString("track")
	if trackID == "" {
		c.Redirect("http://pl.leadernet-hk.com/identify", 302)
		return
	}
	c.Data["track"] = trackID
	c.TplName = "dm/pl/lp.html"
}

type PLWelcomePageControllers struct {
	beego.Controller
}
