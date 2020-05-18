package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/models"
)

type PlanController struct {
	beego.Controller
}

func (c *PlanController) URLMapping() {
	c.Mapping("GetAllReceiver", c.GetAllReceiver)
	c.Mapping("AddReceiver", c.AddReceiver)
	c.Mapping("GetAllPlans", c.GetAllPlans)
	c.Mapping("AddPlan", c.AddPlan)
	c.Mapping("UpdatePlan", c.UpdatePlan)
	c.Mapping("DeletePlan", c.DeletePlan)
}

// @router / [get]
func (c *PlanController) GetAllPlans() {
	var Receiver *models.Plans
	plans := Receiver.GetAllPlans()
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: plans,
	}
	c.ServeJSON()
}

// @router / [post]
func (c *PlanController) AddPlan() {
	var plan models.Plans
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &plan)
	if err != nil {
		logs.Error("Unmarshal plan error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		err = plan.AddPlan()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, plan)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:planid/receivers/ [get]
func (c *PlanController) GetAllReceiver() {
	planId := c.Ctx.Input.Param(":planid")
	var Receiver *models.Receivers
	receivers := Receiver.GetAllReceivers(planId)
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: receivers,
	}
	c.ServeJSON()
}

// @router /:planid/receivers/ [post]
func (c *PlanController) AddReceiver() {
	planId := c.Ctx.Input.Param(":planid")
	var receiver models.Receivers
	var ans common.Res
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
				id, _ := strconv.ParseInt(planId, 10, 64)
				receiver.Plan = &models.Plans{Id: id}
				err = receiver.AddReceiver()
				if err != nil {
					ans.Code = 1
					ans.Msg = err.Error()
				}
			}
		} else {
			id, _ := strconv.ParseInt(planId, 10, 64)
			receiver.Plan = &models.Plans{Id: id}
			err = receiver.AddReceiver()
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

// @router /:planid [put]
func (c *PlanController) UpdatePlan() {
	var plan models.Plans
	planId := c.Ctx.Input.Param(":planid")
	id, _ := strconv.ParseInt(planId, 10, 64)
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &plan)
	if err == nil {
		plan.Id = id
		err = plan.UpdatePlan()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, plan)
	} else {
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:planid [delete]
func (c *PlanController) DeletePlan() {
	planId := c.Ctx.Input.Param(":planid")
	id, _ := strconv.ParseInt(planId, 10, 64)
	var Receiver *models.Plans
	var ans common.Res
	err := Receiver.DeletePlan(id)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %s", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, planId)
	c.Data["json"] = &ans
	c.ServeJSON()
}
