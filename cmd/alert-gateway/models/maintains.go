package models

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Maintains struct {
	Id        int64  `orm:"auto" json:"id,omitempty"`
	Flag      bool   `json:"flag"`
	TimeStart string `orm:"size(15)" json:"time_start"`
	TimeEnd   string `orm:"size(15)" json:"time_end"`
	Month     int    `json:"month"`
	DayStart  int8   `json:"day_start"`
	DayEnd    int8   `json:"day_end"`
	//Week_start  int8   `json:"week_start"`
	//Week_end    int8   `json:"week_end"`
	//Month_start int8   `json:"month_start"`
	//Month_end   int8   `json:"month_end"`
	Valid *time.Time `json:"valid"`
}

func (*Maintains) TableName() string {
	return "maintain"
}

func (u *Maintains) TableIndex() [][]string {
	return [][]string{
		[]string{"Valid", "DayStart", "DayEnd", "Flag", "TimeStart", "TimeEnd"},
	}
}

func (u *Maintains) GetAllMaintains() interface{} {
	maintains := []Maintains{}
	Ormer().QueryTable(new(Maintains)).Limit(-1).All(&maintains)
	type data struct {
		Id        int64  `json:"id"`
		TimeStart string `json:"time_start"`
		TimeEnd   string `json:"time_end"`
		Month     string `json:"month"`
		DayStart  int8   `json:"day_start"`
		DayEnd    int8   `json:"day_end"`
		Valid     string `json:"valid"`
	}
	res := []data{}
	for _, i := range maintains {
		monthList := []string{}
		for m := 1; m <= 12; m++ {
			if i.Month&int(math.Pow(2, float64(m))) > 0 {
				monthList = append(monthList, strconv.Itoa(m))
			}
		}
		res = append(res, data{
			Id:        i.Id,
			TimeStart: i.TimeStart,
			TimeEnd:   i.TimeEnd,
			Month:     strings.Join(monthList, ","),
			DayStart:  i.DayStart,
			DayEnd:    i.DayEnd,
			Valid:     i.Valid.Format("2006-01-02 15:04:05"),
		})
	}
	return res
}

func (u *Maintains) AddMaintains(hosts string) error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Insert(u)
	if err == nil {
		hostsList := []Hosts{}
		hosts = strings.Replace(hosts, "\r", "", -1)
		for _, i := range strings.Split(hosts, string(10)) {
			if i != "" {
				hostsList = append(hostsList, Hosts{Mid: u.Id, Hostname: i})
			}
		}
		_, err = o.InsertMulti(5000, hostsList)
		if err == nil {
			o.Commit()
		} else {
			o.Rollback()
		}
	} else {
		o.Rollback()
	}
	return errors.Wrap(err, "database insert error")
}

func (u *Maintains) UpdateMaintains(hosts string) error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(u)
	if err == nil {
		_, err = o.Raw("DELETE FROM host WHERE mid = ?", u.Id).Exec()
		if err == nil {
			hostsList := []Hosts{}
			hosts = strings.Replace(hosts, "\r", "", -1)
			for _, i := range strings.Split(hosts, string(10)) {
				if i != "" {
					hostsList = append(hostsList, Hosts{Mid: u.Id, Hostname: i})
				}
			}
			_, err = o.InsertMulti(5000, hostsList)
			if err == nil {
				o.Commit()
			} else {
				o.Rollback()
			}
		} else {
			o.Rollback()
		}
	} else {
		o.Rollback()
	}
	return errors.Wrap(err, "database update error")
}

func (u *Maintains) DeleteMaintains(id string) error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Raw("DELETE FROM maintain WHERE id = ?", id).Exec()
	if err == nil {
		_, err = o.Raw("DELETE FROM host WHERE mid = ?", id).Exec()
		if err == nil {
			o.Commit()
		} else {
			o.Rollback()
		}
	} else {
		o.Rollback()
	}
	return errors.Wrap(err, "database delete error")
}
