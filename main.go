package main

import (
	_ "github.com/MobileCPX/FunGame/initial"
	"github.com/MobileCPX/FunGame/models"
	_ "github.com/MobileCPX/FunGame/routers"
	"github.com/astaxie/beego"
	// "github.com/MobileCPX/FunGame/models"
)

func main() {
	models.FileToSql()
	beego.Run()
}
