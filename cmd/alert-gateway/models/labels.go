package models

//import (
//	"github.com/astaxie/beego/orm"
//)
//
//type Labels struct {
//	Id    int64    `orm:"auto" json:"id,omitempty"`
//	Label string   `orm:"unique;size(255)" json:"label"`
//	Rule  []*Rules `orm:"reverse(many);rel_through(alert-gateway/models.RuleLabels)" json:"rules,omitempty"`
//}
//
//func (*Labels) TableName() string {
//	return "label"
//}
//
//func (g *Labels) GetAll() []Labels {
//	labels := []Labels{}
//	Ormer().QueryTable(new(Labels)).Limit(-1).All(&labels)
//	return labels
//}
//
//func (g *Labels) AddLabel() error {
//	_, err := Ormer().Insert(g)
//	return err
//}
//
//func (g *Labels) DeleteLabel(id string) error {
//	records := []struct {
//		Id int64
//	}{}
//	o := orm.NewOrm()
//	o.Begin()
//	_, err := o.Raw("SELECT id FROM rule_label WHERE label_id=? LOCK IN SHARE MODE", id).QueryRows(&records)
//	if err == nil {
//		if len(records) != 0 {
//			rec:=[]int64{}
//			for _,i:=range records{
//				rec=append(rec,i.Id)
//			}
//			_, err = o.QueryTable("rule_label").Limit(-1).Filter("id__in", rec).Delete()
//			if err != nil {
//				o.Rollback()
//				return err
//			}
//		}
//		_, err = o.Raw("DELETE FROM `label` WHERE id=?", id).Exec()
//		if err == nil {
//			o.Commit()
//			return nil
//		} else {
//			o.Rollback()
//			return err
//		}
//	} else {
//		o.Rollback()
//		return err
//	}
//}
