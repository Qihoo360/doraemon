package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"

	"doraemon/cmd/alert-gateway/common"
	"doraemon/cmd/alert-gateway/logs"
	"doraemon/cmd/alert-gateway/models"
)

type ManageController struct {
	beego.Controller
}

func (c *ManageController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("AddManage", c.AddManage)
	c.Mapping("UpdateManage", c.UpdateManage)
	c.Mapping("DeleteManage", c.DeleteManage)
}

// @router / [get]
func (c *ManageController) GetAll() {
	var manage *models.Manages
	res := manage.GetAllManage()
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: res,
	}
	c.ServeJSON()
}

// @router / [post]
func (c *ManageController) AddManage() {
	var manage models.Manages
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &manage)
	if err != nil {
		logs.Error("Unmarshal prom error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		err = manage.AddManage()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, manage)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [put]
func (c *ManageController) UpdateManage() {
	var manage models.Manages
	var ans common.Res
	manageId := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(manageId, 10, 64)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &manage)
	if err != nil {
		logs.Error("Unmarshal prom error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		manage.Id = id
		err = manage.UpdateManage()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, manage)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [delete]
func (c *ManageController) DeleteManage() {
	manageId := c.Ctx.Input.Param(":id")
	var manage *models.Manages
	var ans common.Res
	err := manage.DeleteManage(manageId)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %s", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, manageId)
	c.Data["json"] = &ans
	c.ServeJSON()
}
