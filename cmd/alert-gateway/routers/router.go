// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	//	"strconv"
	//	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"

	"github.com/Qihoo360/doraemon/cmd/alert-gateway/common"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers"
)

var FilterUser = func(ctx *context.Context) {
	//fmt.Println(ctx.Input.Header("User-Agent"))
	//timestamp, _ := strconv.ParseInt(ctx.Input.Header("Content-Time"), 10, 64)
	//now := time.Now().UnixNano() / int64(time.Millisecond)
	//delta := (now - timestamp) / 1000
	//tmp := strconv.FormatInt(delta,10)
	//if (delta < -60 || delta > 60) && (len(ctx.Request.RequestURI) < 27 || ctx.Request.RequestURI[:27] != "/api/v1/login/oauthcallback") {
	//	_ = ctx.Output.JSON(common.Res{Code: -1, Msg: "Expired requests"}, false, false)
	//}
	//ctx.Output.JSON(struct {
	//	Error string  `json:"error"`
	//}{"invalid request"},false,false)
	//fmt.Println(ctx.Request.Method)
	req := ctx.Request
	requestURI := req.RequestURI
	method := req.Method

	if ((method == "GET" && len(requestURI) >= 13 && (requestURI[:13] == "/api/v1/proms" || requestURI[:13] == "/api/v1/rules")) ||
		(method == "POST" && len(requestURI) >= 14 && requestURI[:14] == "/api/v1/alerts")) &&
		ctx.Input.Header("Token") == "96smhbNpRguoJOCEKNrMqQ" {
		return
	}
	if len(requestURI) >= 14 && requestURI[:14] == "/api/v1/logout" {
		return
	}
	username, _ := ctx.Input.Session("username").(string)
	if username == "" && ctx.Request.RequestURI[:13] != "/api/v1/login" {
		_ = ctx.Output.JSON(common.Res{Code: -1, Msg: "Unauthorized"}, false, false)
	}
}

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//AllowAllOrigins: true,
		AllowOrigins:     []string{"http://10.*.*.*:*", "http://localhost:*", "http://127.0.0.1:*", "http://172.*.*.*:*", "http://192.*.*.*:*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*", "content-time"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("/api/v1/*", beego.BeforeRouter, FilterUser)

	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/logout",
			beego.NSInclude(
				&controllers.LogoutController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/rules",
			beego.NSInclude(
				&controllers.RuleController{},
			),
		),
		beego.NSNamespace("/alerts",
			beego.NSInclude(
				&controllers.AlertController{},
			),
		),
		beego.NSNamespace("/plans",
			beego.NSInclude(
				&controllers.PlanController{},
			),
		),
		beego.NSNamespace("/receivers",
			beego.NSInclude(
				&controllers.ReceiverController{},
			),
		),
		beego.NSNamespace("/groups",
			beego.NSInclude(
				&controllers.GroupController{},
			),
		),
		beego.NSNamespace("/proms",
			beego.NSInclude(
				&controllers.PromController{},
			),
		),
		beego.NSNamespace("/maintains",
			beego.NSInclude(
				&controllers.MaintainController{},
			),
		),
		beego.NSNamespace("/manages",
			beego.NSInclude(
				&controllers.ManageController{},
			),
		),
		beego.NSNamespace("/configs",
			beego.NSInclude(
				&controllers.ConfigController{},
			),
		),
		//beego.NSNamespace("/labels",
		//	beego.NSInclude(
		//		&controllers.LabelController{},
		//	),
		//),
	)
	beego.AddNamespace(ns)
}
