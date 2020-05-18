package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Manages struct {
	Id          int64  `orm:"auto" json:"id,omitempty"`
	ServiceName string `orm:"column(servicename);unique;size(255)" json:"servicename"`
	Type        string `orm:"size(255)" json:"type"`
	Status      int8   `orm:"index" json:"status"`
}

func (*Manages) TableName() string {
	return "manage"
}

func (p *Manages) GetAllManage() []Manages {
	manages := []Manages{}
	Ormer().QueryTable(new(Manages)).Limit(-1).All(&manages)
	return manages
}

func (p *Manages) AddManage() error {
	_, err := Ormer().Insert(p)
	return errors.Wrap(err, "database insert error")
}

func (p *Manages) UpdateManage() error {
	_, err := Ormer().Update(p)
	return errors.Wrap(err, "database update error")
}

func (p *Manages) DeleteManage(id string) error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Raw("DELETE FROM manage WHERE id = ?", id).Exec()
	if err == nil {
		_, err = o.Raw("DELETE FROM config WHERE service_id = ?", id).Exec()
		if err != nil {
			o.Rollback()
			return errors.Wrap(err, "database delete error")
		}
	} else {
		o.Rollback()
		return errors.Wrap(err, "database delete error")
	}
	o.Commit()
	return errors.Wrap(err, "database delete error")
}
