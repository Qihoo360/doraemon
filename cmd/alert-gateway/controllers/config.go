package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"

	"doraemon/cmd/alert-gateway/common"
	"doraemon/cmd/alert-gateway/logs"
	"doraemon/cmd/alert-gateway/models"
)

type ConfigController struct {
	beego.Controller
}

func (c *ConfigController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("AddConfig", c.AddConfig)
	c.Mapping("UpdateConfig", c.UpdateConfig)
	c.Mapping("DeleteConfig", c.DeleteConfig)
}

// @router / [get]
func (c *ConfigController) GetAll() {
	idc := c.Input().Get("idc")
	var config *models.Configs
	data := config.GetAllConfig(idc)
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: data,
	}
	c.ServeJSON()
}

// @router / [post]
func (c *ConfigController) AddConfig() {
	var config *models.Configs
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &config)
	if err != nil {
		logs.Error("Unmarshal prom error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		err = config.AddConfig()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, config)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [put]
func (c *ConfigController) UpdateConfig() {
	var config *models.Configs
	var ans common.Res
	configId := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(configId, 10, 64)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &config)
	if err != nil {
		logs.Error("Unmarshal prom error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		config.Id = id
		err = config.UpdateConfig()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, config)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [delete]
func (c *ConfigController) DeleteConfig() {
	configId := c.Ctx.Input.Param(":id")
	var config *models.Configs
	var ans common.Res
	err := config.DeleteConfig(configId)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %s", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, configId)
	c.Data["json"] = &ans
	c.ServeJSON()
}
