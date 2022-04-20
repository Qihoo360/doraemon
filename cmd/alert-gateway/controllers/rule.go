package controllers

import (
	"encoding/json"
	"runtime"
	"strconv"

	"github.com/astaxie/beego"

	"doraemon/cmd/alert-gateway/common"
	"doraemon/cmd/alert-gateway/logs"
	"doraemon/cmd/alert-gateway/models"
)

type RuleController struct {
	beego.Controller
}

func (c *RuleController) URLMapping() {
	c.Mapping("GetRules", c.GetRules)
	c.Mapping("GetAllRules", c.GetAllRules)
	c.Mapping("AddRule", c.AddRule)
	c.Mapping("UpdateRule", c.UpdateRule)
	c.Mapping("DeleteRule", c.DeleteRule)
}

type Rule struct {
	Id          int64  `json:"id"`
	Expr        string `json:"expr"`
	Op          string `json:"op"`
	Value       string `json:"value"`
	For         string `json:"for"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	PromId      int64  `json:"prom_id"`
	PlanId      int64  `json:"plan_id"`
}

var rule struct {
	Expr        string `json:"expr"`
	For         string `json:"for"`
	Op          string `json:"op"`
	Value       string `json:"value"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	PromId      int64  `json:"prom_id"`
	PlanId      int64  `json:"plan_id"`
}

// @router /getall [get]
func (c *RuleController) GetAllRules() {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in GetAllRules:%v\n%s", e, buf)
		}
	}()

	prom := c.Input().Get("prom")
	id := c.Input().Get("id")
	var Receiver *models.Rules
	rules := Receiver.Get(prom, id)
	res := []Rule{}
	for _, i := range rules {

		res = append(res, Rule{
			Id:          i.Id,
			Expr:        i.Expr,
			Op:          i.Op,
			Value:       i.Value,
			For:         i.For,
			Summary:     i.Summary,
			Description: i.Description,
			PromId:      i.PromId,
			PlanId:      i.PlanId,
		})
	}

	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: res,
	}
	c.ServeJSON()
}

// @router / [get]
func (c *RuleController) GetRules() {
	pageNo, _ := strconv.ParseInt(c.Input().Get("page"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.Input().Get("pagesize"), 10, 64)
	if pageNo == 0 && pageSize == 0 {
		pageNo = 1
		pageSize = 20
	}
	var Receiver *models.Rules
	rules := Receiver.GetRules(pageNo, pageSize)
	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: rules,
	}
	c.ServeJSON()
}

// @router / [post]
func (c *RuleController) AddRule() {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in AddRule:%v\n%s", e, buf)
		}
	}()
	var ruleModel models.Rules
	var ans common.Res

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ruleModel)
	if err != nil {
		logs.Error("Unmarshal rule error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		ruleModel.Id = 0 //reset the "Id" to 0,which is very important:after a record is inserted,the value of "Id" will not be 0,but the auto primary key of the record

		err = ruleModel.InsertRule()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.Ctx.Request.RequestURI, c.Ctx.Request.Method, rule)
	}

	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:ruleid [put]
func (c *RuleController) UpdateRule() {
	ruleId := c.Ctx.Input.Param(":ruleid")
	var ruleModel models.Rules
	var rule struct {
		Expr        string `json:"expr"`
		Op          string `json:"op"`
		Value       string `json:"value"`
		For         string `json:"for"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		PromId      int64  `json:"prom_id"`
		PlanId      int64  `json:"plan_id"`
	}
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &rule)
	if err != nil {
		logs.Error("Unmarshal rule error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		id, _ := strconv.ParseInt(ruleId, 10, 64)
		ruleModel.Id = id
		ruleModel.Expr = rule.Expr
		ruleModel.Op = rule.Op
		ruleModel.Value = rule.Value
		ruleModel.For = rule.For
		ruleModel.Description = rule.Description
		ruleModel.Summary = rule.Summary
		ruleModel.PromId = rule.PromId
		ruleModel.PlanId = rule.PlanId

		err = ruleModel.UpdateRule()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, ruleId)
	}
	c.Data["json"] = &ans
	c.ServeJSON()
}

// @router /:ruleid [delete]
func (c *RuleController) DeleteRule() {
	ruleId := c.Ctx.Input.Param(":ruleid")
	var Receiver *models.Rules
	var ans common.Res
	err := Receiver.DeleteRule(ruleId)
	if err != nil {
		ans.Code = 1
		ans.Msg = err.Error()
	}
	logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, ruleId)
	c.Data["json"] = &ans
	c.ServeJSON()
}
