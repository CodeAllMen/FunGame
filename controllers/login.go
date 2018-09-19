package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	switch c.Ctx.Input.Param(":method") {
	case "login":
		if c.GetString("user") == "admin" && c.GetString("pass") == "test123" {
			c.SetSession("user", c.GetString("user"))
		}
	}
	c.Redirect("/", 302)
}
