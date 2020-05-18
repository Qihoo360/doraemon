package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Plans struct {
	Id          int64  `orm:"auto" json:"id,omitempty"`
	RuleLabels  string `orm:"column(rule_labels);size(255)" json:"rule_labels"`
	Description string `orm:"column(description);size(1023)" json:"description"`
}

func (*Plans) TableName() string {
	return "plan"
}

func (plan *Plans) GetAllPlans() []Plans {
	plans := []Plans{}
	Ormer().QueryTable(new(Plans)).Limit(-1).All(&plans)
	return plans
}

func (plan *Plans) AddPlan() error {
	_, err := Ormer().Insert(plan)
	return errors.Wrap(err, "database insert error")
}

func (plan *Plans) UpdatePlan() error {
	_, err := Ormer().Update(plan)
	return errors.Wrap(err, "database update error")
}

func (plan *Plans) DeletePlan(id int64) error {
	var rules []struct{ Id int64 }
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Raw("SELECT id FROM rule WHERE plan_id = ? LOCK IN SHARE MODE", id).QueryRows(&rules)
	if err == nil {
		if len(rules) > 0 {
			o.Commit()
			return fmt.Errorf("cannot delete this plan,it is associated with following rules:%v", rules)
		} else {
			_, err = o.Raw("DELETE FROM plan WHERE id = ?", id).Exec()
			if err == nil {
				_, err = o.Raw("DELETE FROM plan_receiver WHERE plan_id = ?", id).Exec()
				if err != nil {
					o.Rollback()
					return errors.Wrap(err, "database delete error")
				}
			} else {
				o.Rollback()
				return errors.Wrap(err, "database delete error")
			}
		}
	} else {
		o.Rollback()
		return errors.Wrap(err, "database query error")
	}
	o.Commit()
	return errors.Wrap(err, "database delete error")
}
