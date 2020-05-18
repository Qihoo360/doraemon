package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/models"
)

type ReceiverController struct {
	beego.Controller
}

func (c *ReceiverController) URLMapping() {
	c.Mapping("UpdateReceiver", c.UpdateReceiver)
	c.Mapping("DeleteReceiver", c.DeleteReceiver)
}

// @router /:receiverid [put]
func (c *ReceiverController) UpdateReceiver() {
	var receiver models.Receivers
	var ans common.Res
	receiverId := c.Ctx.Input.Param(":receiverid")
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &receiver)
	if err != nil {
		logs.Error("Unmarshal rule error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		if receiver.Expression != "" {
			root, err := common.BuildTree(receiver.Expression)
			if err != nil {
				ans.Code = 1
				ans.Msg = err.Error()
			} else {
				ReversePolishNotation := common.Converse2ReversePolishNotation(root)
				receiver.ReversePolishNotation = ReversePolishNotation
				id, _ := strconv.ParseInt(receiverId, 10, 64)
				receiver.Id = id
				err = receiver.UpdateReceiver()
				if err != nil {
					ans.Code = 1
					ans.Msg = err.Error()
				}
			}
		} else {
			id, _ := strconv.ParseInt(receiverId, 10, 64)
			receiver.Id = id
			err = receiver.UpdateReceiver()
			if err != nil {
				ans.Code = 1
				ans.Msg = err.Error()
			}
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, receiver)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:receiverid [delete]
func (c *ReceiverController) DeleteReceiver() {
	receiverId := c.Ctx.Input.Param(":receiverid")
	var Receiver *models.Receivers
	var ans common.Res
	err := Receiver.DeleteReceiver(receiverId)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %s", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, receiverId)
	c.Data["json"] = &ans
	c.ServeJSON()
}
