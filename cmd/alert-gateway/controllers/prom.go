package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/models"
)

type PromController struct {
	beego.Controller
}

func (c *PromController) URLMapping() {
	c.Mapping("GetAllProms", c.GetAllProms)
	c.Mapping("AddProm", c.AddProm)
	c.Mapping("UpdateProm", c.UpdateProm)
	c.Mapping("DeleteProm", c.DeleteProm)
}

// @router / [get]
func (c *PromController) GetAllProms() {
	var Receiver *models.Proms
	proms := Receiver.GetAllProms()
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: proms,
	}
	c.ServeJSON()
}

// @router / [post]
func (c *PromController) AddProm() {
	var prom models.Proms
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &prom)
	if err != nil {
		logs.Error("Unmarshal prom error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		err = prom.AddProms()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, prom)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [put]
func (c *PromController) UpdateProm() {
	var prom models.Proms
	var ans common.Res
	promId := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(promId, 10, 64)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &prom)
	if err == nil {
		prom.Id = id
		err = prom.UpdateProms()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, prom)
	} else {
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [delete]
func (c *PromController) DeleteProm() {
	promId := c.Ctx.Input.Param(":id")
	var Receiver *models.Proms
	var ans common.Res
	err := Receiver.DeleteProms(promId)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %s", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, promId)
	c.Data["json"] = &ans
	c.ServeJSON()
}
