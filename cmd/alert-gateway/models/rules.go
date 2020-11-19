package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
)

type Rules struct {
	Id          int64  `orm:"column(id);auto" json:"id,omitempty"`
	Expr        string `orm:"column(expr);size(1023)" json:"expr"`
	Op          string `orm:"column(op);size(31)" json:"op"`
	Value       string `orm:"column(value);size(1023)" json:"value"`
	For         string `orm:"column(for);size(1023)" json:"for"`
	Summary     string `orm:"column(summary);size(1023)" json:"summary"`
	Description string `orm:"column(description);size(1023)" json:"description"`
	PromId      	int64  `orm:"column(prom_id);" json:"prom_id"`
	PlanId      	int64  `orm:"column(plan_id);" json:"plan_id"`
	//Prom        	*Proms `orm:"rel(fk)" json:"prom_id"`
	//Plan        	*Plans `orm:"rel(fk)" json:"plan_id"`
	//Labels      []*Labels `orm:"rel(m2m);rel_through(alert-gateway/models.RuleLabels)" json:"omitempty"`
}

type ShowRules struct {
	Rules []Rules	 `json:"rows"`
	Total  int64     `json:"total"`
}

func (*Rules) TableName() string {
	return "rule"
}

func (rule *Rules) DeleteRule(id string) error {
	_, err := Ormer().Raw("DELETE FROM rule WHERE id = ?", id).Exec()
	return errors.Wrap(err, "database delete error")
}

func (rule *Rules) UpdateRule() error {
	var prom []struct{ PromId int64 }
	var plan []struct{ PlanId int64 }
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Raw("SELECT id FROM prom WHERE id = ? LOCK IN SHARE MODE", rule.PromId).QueryRows(&prom)
	if (err != nil) || (len(prom) == 0) {
		o.Rollback()
		logs.Error("The prom_id %s is invalid", rule.PromId)
		return fmt.Errorf("invalid prom_id %v", rule.PromId)
	}
	_, err = o.Raw("SELECT id FROM plan WHERE id = ? LOCK IN SHARE MODE", rule.PlanId).QueryRows(&plan)
	if (err != nil) || (len(plan) == 0) {
		o.Rollback()
		logs.Error("The plan_id %s is invalid", rule.PlanId)
		return fmt.Errorf("invalid plan_id %v", rule.PlanId)
	}

	_, err = o.Update(rule)
	if err != nil {
		o.Rollback()
		logs.Error("update rule error:%v", err)
		return errors.Wrap(err, "database update error")
	}
	o.Commit()

	return errors.Wrap(err, "database insert error")
}

func (rule *Rules) InsertRule() error {
	var prom []struct{ PromId int64 }
	var plan []struct{ PlanId int64 }
	o := orm.NewOrm()

	o.Begin()
	_, err := o.Raw("SELECT id FROM prom WHERE id = ? LOCK IN SHARE MODE", rule.PromId).QueryRows(&prom)
	if (err != nil) || (len(prom) == 0) {
		o.Rollback()
		logs.Error("The prom_id %s is invalid", rule.PromId)
		return fmt.Errorf("invalid prom_id %v", rule.PromId)
	}
	_, err = o.Raw("SELECT id FROM plan WHERE id = ? LOCK IN SHARE MODE", rule.PlanId).QueryRows(&plan)
	if (err != nil) || (len(plan) == 0) {
		o.Rollback()
		logs.Error("The plan_id %s is invalid", rule.PlanId)
		return fmt.Errorf("invalid plan_id %v", rule.PlanId)
	}

	_, err = o.Insert(rule)
	if err != nil {
		o.Rollback()
		logs.Error("Insert rule error:%v", err)
		return errors.Wrap(err, "database insert error")
	}
	o.Commit()

	return errors.Wrap(err, "database insert error")
}

func (*Rules) Get(prom string, id string) []Rules {
	rules := []Rules{}
	qs := Ormer().QueryTable(new(Rules))
	cond := orm.NewCondition()
	if prom != "" {
		qs = qs.SetCond(cond.And("profile__prom_id__eq", prom))
	} else if id != ""{
		qs = qs.SetCond(cond.And("profile__id__eq", id))
	}

	qs.Limit(-1).All(&rules)
	return rules
}


func (*Rules) GetRules(pageNo int64, pageSize int64) ShowRules{
	var showRules ShowRules
	qs := Ormer().QueryTable(new(Rules))

	// 处理完查询条件之后
	showRules.Total, _ = qs.Count()
	qs.Limit(pageSize).Offset((pageNo-1)*pageSize).All(&showRules.Rules)

	return showRules
}