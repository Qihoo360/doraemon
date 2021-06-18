package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
)

type Inhibits struct {
	Id            			int64      `orm:"column(id);auto" json:"id,omitempty"`
	Name         			string     `orm:"column(name);size(255)" json:"name"`
	SourceExpression        string     `orm:"column(source_expression);size(1023)" json:"source_expression"`
	SourceReversePolishNotation        	string     `orm:"column(source_reverse_polish_notation);size(1023)" json:"source_reverse_polish_notation"`
	Targetxpression        				string     `orm:"column(target_expression);size(1023)" json:"target_expression"`
	TargetReversePolishNotation        	string     `orm:"column(target_reverse_polish_notation);size(1023)" json:"target_reverse_polish_notation"`
	Labels        			string     `orm:"column(labels);size(1023)" json:"labels"`
}

type ShowInhibits struct {
	Inhibits []Inhibits	 `json:"rows"`
	Total  int64     	`json:"total"`
}

func (*Inhibits) TableName() string {
	return "inhibits"
}

func (inhibits *Inhibits) DeleteInhibit(id int64) error {
	if _, err := orm.NewOrm().Delete(&Inhibits{Id: id}); err != nil {
		return errors.Wrap(err, "database delete error")
	}
	return errors.Wrap(nil, "success")
}

func (inhibits *Inhibits) UpdateInhibit() error {
	o := orm.NewOrm()
	_, err := o.Update(inhibits)
	if err != nil {
		logs.Error("update inhibits error:%v", err)
		return errors.Wrap(err, "database update error")
	}
	return errors.Wrap(err, "database insert error")
}

func (inhibits *Inhibits) InsertInhibit() error {
	o := orm.NewOrm()
	_, err := o.Insert(inhibits)
	if err != nil {
		logs.Error("Insert inhibits error:%v", err)
		return errors.Wrap(err, "database insert error")
	}
	return errors.Wrap(err, "database insert error")
}

func (*Inhibits) Get(id int64) Inhibits{
	var inhibit Inhibits
	Ormer().QueryTable(new(Inhibits)).Filter("id__eq", id).One(&inhibit)
	return inhibit
}

func (*Inhibits) GetInhibits(pageNo int64, pageSize int64) ShowInhibits{
	var showInhibits ShowInhibits
	qs := Ormer().QueryTable(new(Inhibits))

	// 处理完查询条件之后
	showInhibits.Total, _ = qs.Count()
	qs.Limit(pageSize).Offset((pageNo-1)*pageSize).All(&showInhibits.Inhibits)

	return showInhibits
}