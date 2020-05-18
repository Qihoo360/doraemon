package main

import (
	"github.com/astaxie/beego"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/initial"
	_ "github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	_ "github.com/Qihoo360/doraemon/cmd/alert-gateway/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	initial.InitDb()
	beego.Run()
}
