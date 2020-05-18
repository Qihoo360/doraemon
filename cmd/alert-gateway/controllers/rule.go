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

type RuleController struct {
	beego.Controller
}

func (c *RuleController) URLMapping() {
	//c.Mapping("GetAllRules", c.GetAllRules)
	c.Mapping("SendAllRules", c.SendAllRules)
	c.Mapping("AddRule", c.AddRule)
	c.Mapping("UpdateRule", c.UpdateRule)
	c.Mapping("DeleteRule", c.DeleteRule)
}

type Rule struct {
	Id    int64  `json:"id"`
	Expr  string `json:"expr"`
	Op    string `json:"op"`
	Value string `json:"value"`
	For   string `json:"for"`
	//Labels      map[string]string `json:"labels"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	PromId      int64  `json:"prom_id"`
	PlanId      int64  `json:"plan_id"`
}

// @router / [get]
func (c *RuleController) SendAllRules() {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 16384)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Panic.Error("Panic in SendAllRules:%v\n%s", e, buf)
		}
	}()
	//prom := c.Input().Get("prom")
	//id:=c.Input().Get("id")
	//summary:=c.Input().Get("summary")
	//pageNo, _ := strconv.ParseInt(c.Input().Get("page"), 10, 64)
	//pageSize, _ := strconv.ParseInt(c.Input().Get("pagesize"), 10, 64)
	//if pageNo==0{
	//	pageNo=1
	//}
	//if pageSize==0{
	//	pageSize=10
	//}
	//var Receiver *models.Rules
	//rules := Receiver.Get(prom,id,summary,pageNo, pageSize)
	//res := []rule{}
	//for _, i := range rules {
	//	labels := []models.Label{}
	//	l:=map[string] string{}
	//	models.Ormer().Raw("SELECT rule_id,lable_id,value FROM rule_label WHERE rule_id=?", i.Id, &labels)
	//	for _, j := range labels {
	//		l[j.Label]=j.Value
	//	}
	//	res = append(res, rule{i.Id, i.Expr, i.Op, i.Value, i.For, l,i.Summary, i.Description, i.Prom.Id, i.Plan.Id})
	//}

	prom := c.Input().Get("prom")
	id := c.Input().Get("id")
	var Receiver *models.Rules
	rules := Receiver.Get(prom, id)
	res := []Rule{}
	for _, i := range rules {
		//labels := []models.Label{}
		//l:=map[string] string{}
		//models.Ormer().Raw("SELECT rule_id,lable_id,value FROM rule_label WHERE rule_id=?", i.Id, &labels)
		//for _, j := range labels {
		//	l[j.Label]=j.Value
		//}
		res = append(res, Rule{
			Id:          i.Id,
			Expr:        i.Expr,
			Op:          i.Op,
			Value:       i.Value,
			For:         i.For,
			Summary:     i.Summary,
			Description: i.Description,
			PromId:      i.Prom.Id,
			PlanId:      i.Plan.Id,
		})
	}

	c.Data["json"] = &common.Res{
		Code: 0,
		Msg:  "",
		Data: res,
	}
	c.ServeJSON()
}

//// @router / [get]
//func (c *RuleController) GetAllRules() {
//	var Receiver *models.Rules
//	rules := Receiver.GetAll()
//	var res []R2
//	for _, element := range rules {
//		res = append(res, R2{element.Id, element.Expr, element.For, element.Labels, element.Summary, element.Description, element.Prom.Id, element.Plan.Id})
//	}
//	c.Data["json"] = &common.Res{0, "", res}
//	c.ServeJSON()
//}

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
	var ans common.Res
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &rule)
	if err != nil {
		logs.Error("Unmarshal rule error:%v", err)
		ans.Code = 1
		ans.Msg = "Unmarshal error"
	} else {
		ruleModel.Id = 0 //reset the "Id" to 0,which is very important:after a record is inserted,the value of "Id" will not be 0,but the auto primary key of the record
		ruleModel.Expr = rule.Expr
		ruleModel.Op = rule.Op
		ruleModel.Value = rule.Value
		ruleModel.For = rule.For
		ruleModel.Summary = rule.Summary
		ruleModel.Description = rule.Description
		ruleModel.Prom = &models.Proms{Id: rule.PromId} //cannot be models.RulesModel.Prom.Id=1,because the "Prom" is a pointer,which refers the null(cannot dereference the null pointer )
		ruleModel.Plan = &models.Plans{Id: rule.PlanId}
		err = ruleModel.InsertRule()
		if err != nil {
			ans.Code = 1
			ans.Msg = err.Error()
		}
		logs.Logger.Info("%s %s %s %v", c.GetSession("username"), c.Ctx.Request.RequestURI, c.Ctx.Request.Method, rule)
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
		ruleModel.Prom = &models.Proms{Id: rule.PromId}
		ruleModel.Plan = &models.Plans{Id: rule.PlanId}
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

//// @router /:ruleid/labels/ [get]
//func (c *RuleController) GetLabel() {
//	ruleid := c.Ctx.Input.Param(":ruleid")
//	label := models.Rules{}
//	var ans common.Res
//	data := label.GetLabel(ruleid)
//	ans.Data = data
//	c.Data["json"] = &ans
//	c.ServeJSON()
//}

//// @router /:ruleid/labels/:labelid [delete]
//func (c *RuleController) DeleteLabel() {
//	ruleid := c.Ctx.Input.Param(":ruleid")
//	labelid := c.Ctx.Input.Param(":labelid")
//	var label *models.RuleLabels
//	var ans common.Res
//	err := label.DeleteLabel(ruleid, labelid)
//	if err != nil {
//		ans.Code = 1
//		ans.Msg = "数据库删除记录错误：" + err.Error()
//	}
//	c.Data["json"] = &ans
//	c.ServeJSON()
//}

//// @router /:ruleid/labels/:labelid [post]
//func (c *RuleController) AddLabel() {
//	ruleid := c.Ctx.Input.Param(":ruleid")
//	labelid := c.Ctx.Input.Param(":labelid")
//	label := models.RuleLabels{}
//	value := struct {
//		Value string `json:"value"`
//	}{}
//	err := json.Unmarshal(c.Ctx.Input.RequestBody, &value)
//	if err == nil {
//		lid, _ := strconv.ParseInt(labelid, 10, 64)
//		rid, _ := strconv.ParseInt(ruleid, 10, 64)
//		label.LabelId = &models.Labels{Id: lid}
//		label.RuleId = &models.Rules{Id: rid}
//		label.Value = value.Value
//		err = label.AddRuleLabel()
//	}
//	var ans common.Res
//	if err != nil {
//		ans.Code = 1
//		ans.Msg = "添加label错误：" + err.Error()
//	}
//	c.Data["json"] = &ans
//	c.ServeJSON()
//}

//// @router /:ruleid/labels/:labelid [put]
//func (c *RuleController) UpdateLabel() {
//	ruleid := c.Ctx.Input.Param(":ruleid")
//	labelid := c.Ctx.Input.Param(":labelid")
//	label := models.RuleLabels{}
//	value := struct {
//		Value string
//	}{}
//	err := json.Unmarshal(c.Ctx.Input.RequestBody, &value)
//	if err == nil {
//		lid, _ := strconv.ParseInt(labelid, 10, 64)
//		rid, _ := strconv.ParseInt(ruleid, 10, 64)
//		label.LabelId = &models.Labels{Id: lid}
//		label.RuleId = &models.Rules{Id: rid}
//		label.Value = value.Value
//		err = label.UpdateLabel()
//	}
//	var ans common.Res
//	if err != nil {
//		ans.Code = 1
//		ans.Msg = "数据库更新记录错误：" + err.Error()
//	}
//	c.Data["json"] = &ans
//	c.ServeJSON()
//}
