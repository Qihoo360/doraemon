package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Proms struct {
	Id   int64  `orm:"auto" json:"id,omitempty"`
	Name string `orm:"size(1023)" json:"name"`
	Url  string `orm:"size(1023)" json:"url"`
}

func (*Proms) TableName() string {
	return "prom"
}

func (p *Proms) GetAllProms() []Proms {
	proms := []Proms{}
	Ormer().QueryTable(new(Proms)).Limit(-1).All(&proms)
	return proms
}

func (p *Proms) AddProms() error {
	_, err := Ormer().Insert(p)
	return errors.Wrap(err, "database insert error")
}

func (p *Proms) UpdateProms() error {
	_, err := Ormer().Update(p)
	return errors.Wrap(err, "database update error")
}

func (p *Proms) DeleteProms(id string) error {
	var rules []struct{ Id int64 }
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Raw("SELECT id FROM rule WHERE prom_id = ? LOCK IN SHARE MODE", id).QueryRows(&rules)
	if err == nil {
		if len(rules) > 0 {
			o.Commit()
			return fmt.Errorf("cannot delete this record,it is associated with following rules:%v", rules)
		} else {
			_, err = o.Raw("DELETE FROM prom WHERE id = ?", id).Exec()
			if err != nil {
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
