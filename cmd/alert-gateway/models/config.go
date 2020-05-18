package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

var ErrNoService = errors.New("add config failed:the service is not exist,please refresh")

type Configs struct {
	Id        int64  `orm:"auto" json:"id,omitempty"`
	ServiceId int64  `json:"serviceid"`
	Idc       string `orm:"size(255)" json:"idc"`
	Proto     string `orm:"size(255)" json:"proto"`
	Auto      string `orm:"size(255)" json:"auto"`
	Port      int    `json:"port"`
	Metric    string `orm:"size(255)" json:"metric"`
}

func (*Configs) TableName() string {
	return "config"
}

func (p *Configs) GetAllConfig(idc string) []Configs {
	configs := []Configs{}
	if idc != "" {
		Ormer().QueryTable(new(Configs)).Limit(-1).Filter("idc__exact", idc).All(&configs)
	} else {
		Ormer().QueryTable(new(Configs)).Limit(-1).All(&configs)
	}
	return configs
}

func (p *Configs) AddConfig() error {
	var rows []struct{ Id int64 }
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Raw("SELECT id FROM manage WHERE id=? LOCK IN SHARE MODE", p.ServiceId).QueryRows(&rows)
	if err == nil {
		if len(rows) > 0 {
			_, err = o.Insert(p)
			if err == nil {
				o.Commit()
			} else {
				o.Rollback()
			}
		} else {
			o.Commit()
			return ErrNoService
		}
	} else {
		o.Rollback()
		return errors.Wrap(err, "database query error")
	}
	return errors.Wrap(err, "database insert error")
}

func (p *Configs) UpdateConfig() error {
	_, err := Ormer().Update(p, "idc", "proto", "metric", "auto", "port")
	return errors.Wrap(err, "database update error")
}

func (p *Configs) DeleteConfig(id string) error {
	_, err := Ormer().Raw("DELETE FROM config WHERE id = ?", id).Exec()
	return errors.Wrap(err, "database delete error")
}
