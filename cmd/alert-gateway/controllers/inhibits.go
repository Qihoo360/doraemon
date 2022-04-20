package controllers

import (
	"doraemon/cmd/alert-gateway/common"
	"doraemon/cmd/alert-gateway/logs"
	"doraemon/cmd/alert-gateway/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

type InhibitsController struct {
	beego.Controller
}

func (c *InhibitsController) URLMapping() {
	c.Mapping("GetInhibits", c.GetInhibits)
	c.Mapping("GetInhibit", c.GetInhibit)
	c.Mapping("AddInhibit", c.AddInhibit)
	c.Mapping("DeleteInhibit", c.DeleteInhibit)
	c.Mapping("GetInhibitLogs", c.GetInhibitLogs)
}

// @router / [get]
func (c *InhibitsController) GetInhibits() {
	pageNo, _ := strconv.ParseInt(c.Input().Get("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.Input().Get("pagesize"), 10, 64)
	if pageNo == 0 && pageSize == 0 {
		pageNo = 1
		pageSize = 20
	}
	var Receiver *models.Inhibits
	inhibits := Receiver.GetInhibits(pageNo, pageSize)
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: inhibits,
	}
	c.ServeJSON()
}

// @router /:id [get]
func (c *InhibitsController) GetInhibit() {
	var ans common.Res
	inhibitId, err := c.GetInt64("id", 1)
	if err != nil {
		ans.Code = 1
		ans.Msg = "获取参数错误：" + err.Error()
	} else {
		var Receiver *models.Inhibits
		inhibit := Receiver.Get(inhibitId)
		if err != nil {
			ans.Code = 1
			ans.Msg = "数据库删除记录错误：" + err.Error()
		} else {
			ans.Data = inhibit
		}
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router / [post]
func (c *InhibitsController) AddInhibit() {
	inhibit := models.Inhibits{}
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &inhibit)
	if err != nil {
		logs.Error("Unmarshal plan error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		inhibit.Id = 0
		if inhibit.SourceReversePolishNotation == "" {
			inhibit.SourceReversePolishNotation = inhibit.SourceExpression
		}
		if inhibit.TargetReversePolishNotation == "" {
			inhibit.TargetReversePolishNotation = inhibit.Targetxpression
		}
		err = inhibit.InsertInhibit()
		if err != nil {
			ans.Code = 1
			ans.Msg = "插入数据库错误：" + err.Error()
		}
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:id [delete]
func (c *InhibitsController) DeleteInhibit() {
	var ans common.Res
	inhibitId, err := c.GetInt64("id", 1)
	if err != nil {
		ans.Code = 1
		ans.Msg = "获取参数错误：" + err.Error()
	} else {
		var inhibit *models.Inhibits
		err = inhibit.DeleteInhibit(inhibitId)
		if err != nil {
			ans.Code = 1
			ans.Msg = "数据库删除记录错误：" + err.Error()
		}
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /logs [get]
func (c *InhibitsController) GetInhibitLogs() {
	pageNo, _ := strconv.ParseInt(c.Input().Get("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.Input().Get("pagesize"), 10, 64)
	timeStart := c.Input().Get("timestart")
	timeEnd := c.Input().Get("timeend")
	if pageNo == 0 && pageSize == 0 {
		pageNo = 1
		pageSize = 20
	}
	var Receiver *models.InhibitLog
	inhibitLogs := Receiver.GetInhibitLogs(pageNo, pageSize, timeStart, timeEnd)
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: inhibitLogs,
	}
	c.ServeJSON()
}
