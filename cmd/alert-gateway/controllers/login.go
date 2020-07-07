package controllers

import (
	"encoding/json"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"github.com/Qihoo360/doraemon/pkg/auth/ldaputil"
	"github.com/astaxie/beego"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/models"
)

type LoginController struct {
	beego.Controller
}

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type UserInfo struct {
	Name         string `json:"name"`
	Display      string `json:"display"`
	Email        string `json:"email"`
	IsAdmin      bool   `json:"is_admin"`
	AccessToken  string `json:"access_token"`
	ClientID     string `json:"client_id"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

func (c *LoginController) URLMapping() {
	c.Mapping("GetCurrentUser", c.GetCurrentUser)
	c.Mapping("GetMethod", c.GetMethod)
	c.Mapping("Local", c.Local)
}

// @router /method [get]
func (c *LoginController) GetMethod() {
	if c.GetSession("method") == nil {
		c.Data["json"] = &common.Res{
			Code: -1,
			Msg:  "Unauthorized",
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = &common.Res{
			Code: 0,
			Msg:  "",
			Data: c.GetSession("method").(string),
		}
		c.ServeJSON()
	}
}

// @router /username [get]
func (c *LoginController) GetCurrentUser() {
	if c.GetSession("username") == nil {
		c.Data["json"] = &common.Res{
			Code: -1,
			Msg:  "Unauthorized",
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = &common.Res{
			Code: 0,
			Msg:  "",
			Data: c.GetSession("username").(string),
		}
		c.ServeJSON()
	}
}

// @router /local [post]
func (c *LoginController) Local() {
	var auth common.AuthModel
	var res common.Res
	var User *models.Users
	json.Unmarshal(c.Ctx.Input.RequestBody, &auth)
	userInfo, err := User.CheckUser(auth)
	if err == nil {
		c.SetSession("username", userInfo.Username)
		c.SetSession("method", "local")
		res.Msg = "Success"
	} else {
		res.Code = -1
		res.Msg = err.Error()
	}
	c.Data["json"] = &res
	c.ServeJSON()
}

// @router /ldap [post]
func (c *LoginController) Ldap() {
	var auth common.AuthModel
	var res common.Res

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &auth); err != nil {
		res.Code = 1
		res.Msg = "Unmarshal error"
	} else {
		if err := ldaputil.Authenticate(auth.Username, auth.Password); err != nil {
			logs.Info("authenticate fail error: %v", err)
			res.Code = -1
			res.Msg = "Unauthorized"
		} else {
			c.SetSession("username", auth.Username)
			c.SetSession("method", "ldap")
			res.Msg = "Success"
		}
	}

	c.Data["json"] = &res
	c.ServeJSON()
}
