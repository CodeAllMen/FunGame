package main

import (
	_ "github.com/MobileCPX/FunGame/initial"
	_ "github.com/MobileCPX/FunGame/routers"
	"github.com/astaxie/beego"
)

func main() {
	// models.FileToSql()
	beego.Run()
}
