package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/models"
)

type GroupController struct {
	beego.Controller
}

func (c *GroupController) URLMapping() {
	c.Mapping("GetAllGroup", c.GetAllGroup)
	c.Mapping("AddGroup", c.AddGroup)
	c.Mapping("UpdateGroup", c.UpdateGroup)
	c.Mapping("DeleteGroup", c.DeleteGroup)
}

// @router / [get]
func (c *GroupController) GetAllGroup() {
	var Receiver *models.Groups
	groups := Receiver.GetAll()
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: groups,
	}
	c.ServeJSON()
}

// @router / [post]
func (c *GroupController) AddGroup() {
	var group models.Groups
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &group)
	if err != nil {
		logs.Error("Unmarshal plan error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		err = group.AddGroup()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, group)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [put]
func (c *GroupController) UpdateGroup() {
	var group models.Groups
	var ans common.Res
	groupId := c.Ctx.Input.Param(":id")
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &group)
	if err == nil {
		id, _ := strconv.ParseInt(groupId, 10, 64)
		group.Id = id
		err = group.UpdateGroup()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, group)
	} else {
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [delete]
func (c *GroupController) DeleteGroup() {
	groupId := c.Ctx.Input.Param(":id")
	var Receiver *models.Groups
	var ans common.Res
	err := Receiver.DeleteGroup(groupId)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %s", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, groupId)
	c.Data["json"] = &ans
	c.ServeJSON()
}
