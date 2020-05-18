package controllers

import (
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) URLMapping() {
	c.Mapping("Logout", c.Logout)
}

// @router / [get]
func (c *LogoutController) Logout() {
	c.DestroySession()
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "Success",
	}
	c.ServeJSON()
}
