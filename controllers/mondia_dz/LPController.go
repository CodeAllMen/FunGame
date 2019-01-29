package mondia_dz

import "github.com/astaxie/beego"

// mondia 阿尔及利亚  LP页面
type MondiaDZPageController struct {
	beego.Controller
}

// LP首页
func (c *MondiaDZPageController) LpIndex() {
	c.TplName = "mondia/dz/LP.html"
}

// 条款页面
func (c *MondiaDZPageController) Terms() {
	c.TplName = "mondia/dz/terms.html"
}
