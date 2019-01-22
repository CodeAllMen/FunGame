package controllers

import (
	"github.com/MobileCPX/FunGame/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	subID := c.GetString("subID")
	if subID != "" {
		c.SetSession("user", subID)
	}
	c.Data["games"] = models.GetGames()
	c.TplName = "index.html"
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["status"] = "0"
	c.TplName = "login.html"
}

type PageController struct {
	beego.Controller
}

func (c *PageController) Get() {
	sp := c.Ctx.Input.Param(":sp")
	country := c.Ctx.Input.Param(":country")
	pagetype := c.Ctx.Input.Param(":type")

	c.TplName = sp + "/" + country + "/" + pagetype + ".html"
}

type GamePage struct {
	beego.Controller
}

func (c *GamePage) Get() {
	if c.GetSession("user") != nil {
		id := c.Ctx.Input.Param(":id")
		url := models.GetGamesUrl(id)
		c.Redirect(url, 302)
	} else {
		c.TplName = "login.html"
	}
}
