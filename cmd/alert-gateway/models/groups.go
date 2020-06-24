package models

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/pkg/errors"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
)

type Groups struct {
	Id   int64  `orm:"auto" json:"id,omitempty"`
	Name string `orm:"unique;size(255)" json:"name"`
	User string `orm:"size(1023)" json:"user"`
}

type HttpRes struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Mobile  string `json:"mobile"`
		Email   string `json:"email"`
		AddTime string `json:"add_time"`
		Account string `json:"account"`
	} `json:"data"`
}

func (*Groups) TableName() string {
	return "group"
}

func (g *Groups) GetAll() []Groups {
	groups := []Groups{}
	Ormer().QueryTable(new(Groups)).Limit(-1).All(&groups)
	return groups
}

func (g *Groups) AddGroup() error {
	_, err := Ormer().Insert(g)
	return errors.Wrap(err, "database insert error")
}

func (g *Groups) UpdateGroup() error {
	_, err := Ormer().Update(g)
	return errors.Wrap(err, "database update error")
}

func (g *Groups) DeleteGroup(id string) error {
	_, err := Ormer().Raw("DELETE FROM `group` WHERE id=?", id).Exec()
	return errors.Wrap(err, "database delete error")
}

func SendAlertsFor(VUG *common.ValidUserGroup) []string {
	var userList []string
	if VUG.User != "" {
		userList = strings.Split(VUG.User, ",")
	}
	if VUG.Group != "" {
		var groups []*Groups
		_, _ = Ormer().
			QueryTable("group").
			Filter("name__in", strings.Split(VUG.Group, ",")).All(&groups, "user")
		for _, v := range groups {
			userList = append(userList, strings.Split(v.User, ",")...)
		}
	}
	if VUG.DutyGroup != "" {
		date := time.Now().Format("2006-1-2")
		idList := strings.Split(VUG.DutyGroup, ",")
		for _, id := range idList {
			res, _ := common.HttpGet(beego.AppConfig.String("DutyGroupUrl"), map[string]string{"teamId": id, "day": date}, nil)
			info := HttpRes{}
			jsonDataFromHttp, _ := ioutil.ReadAll(res.Body)
			json.Unmarshal(jsonDataFromHttp, &info)
			for _, i := range info.Data {
				userList = append(userList, i.Account)
			}
		}
	}
	hashMap := map[string]bool{}
	for _, name := range userList {
		hashMap[name] = true
	}
	res := []string{}
	for key := range hashMap {
		res = append(res, key)
	}
	return res
}
