package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"doraemon/cmd/alert-gateway/common"
	"doraemon/cmd/alert-gateway/logs"
	"doraemon/cmd/alert-gateway/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("GetAllUser", c.GetAllUser)
	c.Mapping("AddUser", c.AddUser)
	c.Mapping("UpdatePassword", c.UpdatePassword)
	c.Mapping("DeleteUsers", c.DeleteUsers)
}

// @router / [get]
func (c *UserController) GetAllUser() {
	var Receiver *models.Users
	users := Receiver.GetAll()
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: users,
	}
	c.ServeJSON()
}

// @router / [post]
func (c *UserController) AddUser() {
	var userInfo models.Users
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		logs.Error("Unmarshal plan error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		err = userInfo.AddUser()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, userInfo)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router / [put]
func (c *UserController) UpdatePassword() {
	var newInfo struct {
		Name               string `json:"name"`
		OldPassword        string `json:"oldpassword"`
		NewPassword        string `json:"newpassword"`
		ConfirmNewPassword string `json:"confirmnewpassword"`
	}
	var ans common.Res
	var userInfo models.Users
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &newInfo)
	if err == nil {
		if newInfo.ConfirmNewPassword == newInfo.NewPassword {
			if c.GetSession("username") == newInfo.Name {
				err = userInfo.UpdatePassword(newInfo.Name, newInfo.OldPassword, newInfo.NewPassword)
				if err != nil {
					ans.Code = 1
					ans.Msg = err.Error()
				}
			} else {
				ans.Code = 1
				ans.Msg = "Inconsistent user identity"
			}
		} else {
			ans.Code = 1
			ans.Msg = "The two new passwords are inconsistent"
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, newInfo)
	} else {
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [delete]
func (c *UserController) DeleteUsers() {
	userId := c.Ctx.Input.Param(":id")
	var userInfo models.Users
	var ans common.Res
	err := userInfo.DeleteUsers(userId)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %s", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, userId)
	c.Data["json"] = &ans
	c.ServeJSON()
}
