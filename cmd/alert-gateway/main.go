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

//go:generate sh -c "echo 'package routers; import \"github.com/astaxie/beego\"; func init() {beego.BConfig.RunMode = beego.DEV}' > routers/0.go"
//go:generate sh -c "echo 'package routers; import \"os\"; func init() {os.Exit(0)}' > routers/z.go"
//go:generate go run $GOFILE
//go:generate sh -c "rm routers/0.go routers/z.go"