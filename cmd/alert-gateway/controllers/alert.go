package controllers

import (
	"encoding/json"
	"runtime"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/models"
)

type AlertController struct {
	beego.Controller
}

func (c *AlertController) URLMapping() {
	c.Mapping("GetAlerts", c.GetAlerts)
	c.Mapping("ShowAlerts", c.ShowAlerts)
	c.Mapping("Confirm", c.Confirm)
	c.Mapping("HandleAlerts", c.HandleAlerts)
}

// @router / [get]
func (c *AlertController) GetAlerts() {
	pageNo, _ := strconv.ParseInt(c.Input().Get("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.Input().Get("pagesize"), 10, 64)
	timeStart := c.Input().Get("timestart")
	timeEnd := c.Input().Get("timeend")
	status := c.Input().Get("status")
	summary := c.Input().Get("summary")
	if pageNo == 0 && pageSize == 0 {
		pageNo = 1
		pageSize = 10
	}
	var Receiver *models.Alerts
	alerts := Receiver.GetAlerts(pageNo, pageSize, timeStart, timeEnd, status, summary)
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: alerts,
	}
	c.ServeJSON()
}

// @router /rules/:ruleid [get]
func (c *AlertController) ShowAlerts() {
	ruleId := c.Ctx.Input.Param(":ruleid")
	start := c.Input().Get("start")
	pageNo, _ := strconv.ParseInt(c.Input().Get("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.Input().Get("pagesize"), 10, 64)
	var ans common.Res
	var Receiver *models.Alerts
	alerts := Receiver.ShowAlerts(ruleId, start, pageNo, pageSize)
	ans.Data = alerts
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /classify [get]
func (c *AlertController) ClassifyAlerts() {
	var ans common.Res
	var Receiver *models.Alerts
	alerts := Receiver.ClassifyAlerts()
	ans.Data = alerts
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router / [put]
func (c *AlertController) Confirm() {
	var confirmList common.Confirm
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &confirmList)
	if err == nil {
		var Receiver *models.Alerts
		err = Receiver.ConfirmAll(&confirmList)
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, confirmList)
	} else {
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router / [post]
func (c *AlertController) HandleAlerts() {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in HandleAlerts:%v\n%s", e, buf)
		}
	}()
	var alerts common.Alerts
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &alerts)
	logs.Originloger.Info("%v\n", alerts)
	if err != nil {
		logs.Error("Unmarshal error:%s", err)
	} else {
		var Receiver *models.Alerts
		Receiver.AlertsHandler(&alerts)
	}
	var ans common.Res
	c.Data["json"] = &ans
	c.ServeJSON()
}
