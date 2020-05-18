package models

//import (
//	"errors"

//	"github.com/astaxie/beego/orm"
//)
//
//type RuleLabels struct {
//	Id      int64   `orm:"auto" json:"id,omitempty"`
//	RuleId  *Rules  `orm:"rel(fk);column(rule_id)"`
//	LabelId *Labels `orm:"rel(fk);column(label_id)"`
//	Value   string  `orm:"size(255)" json:"value"`
//}
//
//func (*RuleLabels) TableName() string {
//	return "rule_label"
//}
//
//func (p *RuleLabels) TableUnique() [][]string {
//	return [][]string{
//		[]string{"RuleId", "LabelId"},
//	}
//}
//
//func (p *RuleLabels) DeleteLabel(ruleid string, labelid string) error {
//	_, err := Ormer().Raw("DELETE FROM rule_label WHERE rule_id=? AND label_id=?", ruleid, labelid).Exec()
//	return err
//}
//
//func (p *RuleLabels) AddRuleLabel() error {
//	label := []struct {
//		key string
//	}{}
//	o := orm.NewOrm()
//	o.Begin()
//	_, err := o.Raw("SELECT `label` FROM `label` WHERE id=? LOCK IN SHARE MODE", p.LabelId).QueryRows(&label)
//	if err == nil {
//		if len(label) == 0 {
//			o.Commit()
//			return errors.New("label is not exist")
//		} else {
//			_, err = o.Insert(p)
//			if err == nil {
//				o.Commit()
//				return err
//			} else {
//				o.Rollback()
//				return err
//			}
//		}
//	} else {
//		o.Rollback()
//		return err
//	}
//}
//
//func (p *RuleLabels) UpdateLabel() error {
//	_, err := Ormer().Raw("UPDATE rule_label SET value=? WHERE rule_id=? AND label_id=?",p.Value,p.RuleId.Id,p.LabelId.Id).Exec()
//	return err
//}
