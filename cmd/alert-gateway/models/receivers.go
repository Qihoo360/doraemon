package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Receivers struct {
	Id                    int64  `orm:"auto" json:"id,omitempty"`
	Plan                  *Plans `orm:"index;rel(fk)" json:"plan_id"`
	StartTime             string `orm:"size(31)" json:"start_time"`
	EndTime               string `orm:"size(31)" json:"end_time"`
	Start                 int    `json:"start"`
	Period                int    `json:"period"`
	Expression            string `orm:"size(1023)" json:"expression"`
	ReversePolishNotation string `orm:"size(1023)" json:"reverse_polish_notation"`
	User                  string `orm:"size(1023)" json:"user"`
	Group                 string `orm:"size(1023)" json:"group"`
	DutyGroup             string `orm:"size(255)" json:"duty_group"`
	Method                string `orm:"size(255)" json:"method"`
}

type Rec struct {
	Id         int64  `json:"id,omitempty"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Start      int    `json:"start"`
	Period     int    `json:"period"`
	Expression string `json:"expression"`
	User       string `json:"user"`
	Group      string `json:"group"`
	DutyGroup  string `json:"duty_group"`
	Method     string `json:"method"`
}

func (*Receivers) TableName() string {
	return "plan_receiver"
}

func (p *Receivers) GetAllReceivers(planid string) []Rec {
	receivers := []Rec{}
	Ormer().Raw("SELECT id,start_time,end_time,start,period,expression,user,`group`,duty_group,method FROM plan_receiver WHERE plan_id=?", planid).QueryRows(&receivers)
	return receivers
}

func (p *Receivers) AddReceiver() error {
	var planId []struct{ Id int64 }
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Raw("SELECT id FROM plan WHERE id = ? LOCK IN SHARE MODE", p.Plan.Id).QueryRows(&planId)
	if err != nil {
		o.Rollback()
		return errors.Wrap(err, "database query error")
	} else {
		if len(planId) > 0 {
			_, err = o.Insert(p)
			if err != nil {
				o.Rollback()
				return errors.Wrap(err, "database insert error")
			}
		} else {
			o.Commit()
			return fmt.Errorf("plan id: %v is not exsit", p.Plan.Id)
		}
	}
	o.Commit()
	return errors.Wrap(err, "database insert error")
}

func (p *Receivers) UpdateReceiver() error {
	_, err := Ormer().Update(p, "start_time", "end_time", "start", "period", "expression", "reverse_polish_notation", "user", "group", "duty_group", "method")
	return errors.Wrap(err, "database update error")
}

func (p *Receivers) DeleteReceiver(id string) error {
	_, err := Ormer().Raw("DELETE FROM plan_receiver WHERE id=?", id).Exec()
	return errors.Wrap(err, "database delete error")
}
