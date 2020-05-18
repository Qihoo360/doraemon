package models

import (
	"crypto/md5"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
)

type Users struct {
	Id       int64  `orm:"auto" json:"id,omitempty"`
	Name     string `orm:"unique;size(255)" json:"name"`
	Password string `orm:"size(1023)" json:"password,omitempty"`
}

func (*Users) TableName() string {
	return "users"
}

func (g *Users) CheckUser(userInfo common.AuthModel) (*common.AuthModel, error) {
	var queryRes []struct {
		Id       int64
		Password string
	}
	Ormer().Raw("SELECT id,password FROM `users` WHERE name=?", userInfo.Username).QueryRows(&queryRes)
	if len(queryRes) == 0 {
		return nil, errors.Errorf("the user is not exist")
	} else {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(userInfo.Password)))
		if hash != queryRes[0].Password {
			return nil, errors.Errorf("invalid password")
		} else {
			return &userInfo, nil
		}
	}
}

func (g *Users) GetAll() []Users {
	users := []Users{}
	Ormer().QueryTable(new(Users)).Limit(-1).OrderBy("id").All(&users, "id", "name")
	return users
}

func (g *Users) AddUser() error {
	g.Password = fmt.Sprintf("%x", md5.Sum([]byte("123456")))
	_, err := Ormer().Insert(g)
	return errors.Wrap(err, "database insert error")
}

func (g *Users) UpdatePassword(name string, oldPassword string, newPassword string) error {
	var user []Users
	Ormer().Raw("SELECT id,name,password FROM users WHERE name=?", name).QueryRows(&user)
	if len(user) > 0 {
		if user[0].Password == fmt.Sprintf("%x", md5.Sum([]byte(oldPassword))) {
			_, err := Ormer().Raw("UPDATE users SET password=? WHERE name=?", fmt.Sprintf("%x", md5.Sum([]byte(newPassword))), name).Exec()
			return errors.Wrap(err, "database update error")
		} else {
			return errors.Errorf("wrong password")
		}
	} else {
		return errors.Errorf("user is not exist")
	}
}

func (g *Users) DeleteUsers(id string) error {
	_, err := Ormer().Raw("DELETE FROM `users` WHERE id=?", id).Exec()
	return errors.Wrap(err, "database delete error")
}
