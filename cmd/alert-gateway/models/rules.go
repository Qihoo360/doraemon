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
	Value       string `orm:"column(value);size(1023)" json:"op"`
	For         string `orm:"column(for);size(1023)" json:"for"`
	Summary     string `orm:"column(summary);size(1023)" json:"summary"`
	Description string `orm:"column(description);size(1023)" json:"description"`
	Prom        *Proms `orm:"rel(fk)" json:"prom_id"`
	Plan        *Plans `orm:"rel(fk)" json:"plan_id"`
	//Labels      []*Labels `orm:"rel(m2m);rel_through(alert-gateway/models.RuleLabels)" json:"omitempty"`
}

//type Label struct {
//	Labelid int64  `json:"label_id"`
//	Label   string `json:"label"`
//	Value   string `json:"value"`
//}

func (*Rules) TableName() string {
	return "rule"
}

//func (p *Rules) GetLabel(ruleid string) interface{} {
//	Ormer().QueryTable(new(Rules)).Filter("id", ruleid).RelatedSel().One(p)
//	Ormer().LoadRelated(p, "Labels")
//	res := []Label{}
//	for _, i := range p.Labels {
//		v := struct {
//			Value string
//		}{}
//		Ormer().Raw("SELECT value FROM rule_label WHERE label_id=? AND rule_id=?", i.Id, ruleid).QueryRow(&v)
//		res = append(res, Label{i.Id, i.Label, v.Value})
//	}
//	return res
//}

func (rule *Rules) DeleteRule(id string) error {
	_, err := Ormer().Raw("DELETE FROM rule WHERE id = ?", id).Exec()
	//o := orm.NewOrm()
	//o.Begin()
	//_, err := o.Raw("DELETE FROM rule_label WHERE rule_id = ?", id).Exec()
	//if err == nil {
	//	_, err = o.Raw("DELETE FROM rule WHERE id = ?", id).Exec()
	//	if err == nil {
	//		o.Commit()
	//		return err
	//	} else {
	//		o.Rollback()
	//		return err
	//	}
	//} else {
	//	o.Rollback()
	//	return err
	//}
	return errors.Wrap(err, "database delete error")
}

func (rule *Rules) UpdateRule() error {
	var prom []struct{ PromId int64 }
	var plan []struct{ PlanId int64 }
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Raw("SELECT id FROM prom WHERE id = ? LOCK IN SHARE MODE", rule.Prom.Id).QueryRows(&prom)
	//fmt.Println(prom)
	if err == nil && len(prom) > 0 {
		_, err = o.Raw("SELECT id FROM plan WHERE id = ? LOCK IN SHARE MODE", rule.Plan.Id).QueryRows(&plan)
		//fmt.Println(plan)
		if err == nil && len(plan) > 0 {
			_, err = o.Update(rule)
			if err != nil {
				logs.Error("update rule error:%v", err)
				o.Rollback()
				return errors.Wrap(err, "database update error")
			}
		} else {
			o.Rollback()
			logs.Error("The plan_id %s is invalid", rule.Plan.Id)
			return fmt.Errorf("invalid plan_id %v", rule.Plan.Id)
		}
	} else {
		o.Rollback()
		logs.Error("The prom_id %s is invalid", rule.Prom.Id)
		return fmt.Errorf("invalid prom_id %v", rule.Prom.Id)
	}
	o.Commit()
	return errors.Wrap(err, "database update error")
}

func (rule *Rules) InsertRule() error {
	var prom []struct{ PromId int64 }
	var plan []struct{ PlanId int64 }
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Raw("SELECT id FROM prom WHERE id = ? LOCK IN SHARE MODE", rule.Prom.Id).QueryRows(&prom)
	//fmt.Println(prom)
	if err == nil && len(prom) > 0 {
		_, err = o.Raw("SELECT id FROM plan WHERE id = ? LOCK IN SHARE MODE", rule.Plan.Id).QueryRows(&plan)
		//fmt.Println(plan)
		if err == nil && len(plan) > 0 {
			_, err = o.Insert(rule)
			if err != nil {
				logs.Error("Insert rule error:%v", err)
				o.Rollback()
				return errors.Wrap(err, "database insert error")
			}
		} else {
			o.Rollback()
			logs.Error("The plan_id %s is invalid", rule.Plan.Id)
			return fmt.Errorf("invalid plan_id %v", rule.Plan.Id)
		}
	} else {
		o.Rollback()
		logs.Error("The prom_id %s is invalid", rule.Prom.Id)
		return fmt.Errorf("invalid prom_id %v", rule.Prom.Id)
	}
	o.Commit()
	return errors.Wrap(err, "database insert error")
}

func (*Rules) Get(prom string, id string) []Rules {
	//var rules []Rules
	//_, err := Ormer().QueryTable(new(Rules)).Filter("Prom", 1).RelatedSel("Prom").All(&rules)
	//
	//if err == nil {
	//	for _, v := range rules {
	//		fmt.Println("rule:", v)
	//		fmt.Println("级联查询:rule.prom.name:", v.Prom.Name)
	//		fmt.Println("级联查询:rule.prom:", v.Prom)
	//	}
	//}

	//rules := []Rules{}
	//if prom == "" {
	//	if id != "" && summary=="" {
	//		Ormer().QueryTable(new(Rules)).Limit(-1).Filter("id", id).Limit(pageSize,pageNo-1).All(&rules)
	//
	//	}else if id=="" && summary!=""{
	//		Ormer().QueryTable(new(Rules)).Limit(-1).Filter("summary__contains", summary).Limit(pageSize,pageNo-1).All(&rules)
	//	}else {
	//		Ormer().QueryTable(new(Rules)).Limit(-1).Filter("summary__contains",summary).Limit(pageSize,pageNo-1).Filter("id",id).All(&rules)
	//	}
	//} else {
	//	Ormer().QueryTable(new(Rules)).Limit(-1).Filter("prom_id", prom).Limit(pageSize,pageNo-1).All(&rules)
	//}
	//return rules

	rules := []Rules{}
	if prom != "" {
		Ormer().QueryTable(new(Rules)).Limit(-1).Filter("prom_id", prom).All(&rules)
	} else if id != "" {
		Ormer().QueryTable(new(Rules)).Limit(-1).Filter("id", id).All(&rules)
	} else {
		Ormer().QueryTable(new(Rules)).Limit(-1).All(&rules)
	}
	return rules
}
