package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

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
	c.Mapping("OAuthCallback", c.OAuthCallback)
	c.Mapping("OAuthCodeURL", c.OAuthCodeURL)
	c.Mapping("GetCurrentUser", c.GetCurrentUser)
	c.Mapping("GetMethod", c.GetMethod)
	c.Mapping("LDAP", c.LDAP)
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
func (c *LoginController) LDAP() {
	var auth common.AuthModel
	var res common.Res
	json.Unmarshal(c.Ctx.Input.RequestBody, &auth)
	userInfo, err := common.Authenticate(auth)
	if err == nil {
		c.SetSession("username", userInfo.Name)
		c.SetSession("method", "ldap")
		res.Msg = "Success"
	} else {
		res.Code = -1
		res.Msg = err.Error()
	}
	c.Data["json"] = &res
	c.ServeJSON()
}

// @router /oauth [get]
func (c *LoginController) OAuthCodeURL() {
	var res common.Res
	section, err := beego.AppConfig.GetSection("auth.oauth2")
	if err != nil {
		res.Code = -1
		res.Msg = "Can't find OAuth2 config."
		c.ServeJSON()
	} else {
		res.Data = section["auth_url"]
	}
	c.Data["json"] = &res
	c.ServeJSON()
}

// @router /oauthcallback [get]
func (c *LoginController) OAuthCallback() {
	code := c.Input().Get("code")
	if code != "" {
		section, _ := beego.AppConfig.GetSection("auth.oauth2")
		state := c.Input().Get("state")
		decode, _ := url.QueryUnescape(state)
		res, _ := http.PostForm(section["token_url"], url.Values{
			"grant_type":    {"authorization_code"},
			"client_id":     {section["client_id"]},
			"client_secret": {section["client_secret"]},
			"code":          {code},
		})
		jsonDataFromHttp, _ := ioutil.ReadAll(res.Body)
		data := Token{}
		json.Unmarshal(jsonDataFromHttp, &data)
		res, _ = common.HttpGet(section["api_url"], nil, map[string]string{"Authorization": "Bearer " + data.AccessToken})
		jsonDataFromHttp, _ = ioutil.ReadAll(res.Body)
		apiMapping := make(map[string]string)
		userInfo := UserInfo{}
		if section["api_mapping"] != "" {
			for _, km := range strings.Split(section["api_mapping"], ",") {
				arr := strings.Split(km, ":")
				apiMapping[arr[0]] = arr[1]
			}
		}
		if len(apiMapping) == 0 {
			json.Unmarshal(jsonDataFromHttp, &userInfo)
		} else {
			userMap := make(map[string]interface{})
			json.Unmarshal(jsonDataFromHttp, &userMap)
			if userMap[apiMapping["name"]] != nil {
				userInfo.Name = userMap[apiMapping["name"]].(string)
			}
		}
		if userInfo.Name != "" {
			c.SetSession("username", userInfo.Name)
			c.SetSession("method", "oauth")
			c.Redirect(decode, 302)
		} else {
			//logs.Panic.Info("%s", jsonDataFromHttp)
			c.Redirect(decode, 302)
		}
	} else {
		c.Data["json"] = &common.Res{
			Code: -1,
			Msg:  "invalid callback url,lack of parameter code",
		}
		c.ServeJSON()
	}
}
