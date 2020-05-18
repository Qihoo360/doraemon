package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:AlertController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:AlertController"],
		beego.ControllerComments{
			Method:           "GetAlerts",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:AlertController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:AlertController"],
		beego.ControllerComments{
			Method:           "Confirm",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:AlertController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:AlertController"],
		beego.ControllerComments{
			Method:           "HandleAlerts",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:AlertController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:AlertController"],
		beego.ControllerComments{
			Method:           "ClassifyAlerts",
			Router:           `/classify`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:AlertController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:AlertController"],
		beego.ControllerComments{
			Method:           "ShowAlerts",
			Router:           `/rules/:ruleid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ConfigController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ConfigController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ConfigController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ConfigController"],
		beego.ControllerComments{
			Method:           "AddConfig",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ConfigController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ConfigController"],
		beego.ControllerComments{
			Method:           "UpdateConfig",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ConfigController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ConfigController"],
		beego.ControllerComments{
			Method:           "DeleteConfig",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:GroupController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "GetAllGroup",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:GroupController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "AddGroup",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:GroupController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "UpdateGroup",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:GroupController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "DeleteGroup",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "LDAP",
			Router:           `/ldap`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Local",
			Router:           `/local`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "GetMethod",
			Router:           `/method`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "OAuthCodeURL",
			Router:           `/oauth`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "OAuthCallback",
			Router:           `/oauthcallback`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "GetCurrentUser",
			Router:           `/username`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LogoutController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:LogoutController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:MaintainController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:MaintainController"],
		beego.ControllerComments{
			Method:           "GetAllMaintains",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:MaintainController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:MaintainController"],
		beego.ControllerComments{
			Method:           "AddMaintain",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:MaintainController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:MaintainController"],
		beego.ControllerComments{
			Method:           "UpdateMaintain",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:MaintainController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:MaintainController"],
		beego.ControllerComments{
			Method:           "DeleteMaintain",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:MaintainController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:MaintainController"],
		beego.ControllerComments{
			Method:           "GetHosts",
			Router:           `/:id/hosts`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ManageController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ManageController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ManageController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ManageController"],
		beego.ControllerComments{
			Method:           "AddManage",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ManageController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ManageController"],
		beego.ControllerComments{
			Method:           "UpdateManage",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ManageController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ManageController"],
		beego.ControllerComments{
			Method:           "DeleteManage",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"],
		beego.ControllerComments{
			Method:           "GetAllPlans",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"],
		beego.ControllerComments{
			Method:           "AddPlan",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"],
		beego.ControllerComments{
			Method:           "UpdatePlan",
			Router:           `/:planid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"],
		beego.ControllerComments{
			Method:           "DeletePlan",
			Router:           `/:planid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"],
		beego.ControllerComments{
			Method:           "GetAllReceiver",
			Router:           `/:planid/receivers/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PlanController"],
		beego.ControllerComments{
			Method:           "AddReceiver",
			Router:           `/:planid/receivers/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PromController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PromController"],
		beego.ControllerComments{
			Method:           "GetAllProms",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PromController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PromController"],
		beego.ControllerComments{
			Method:           "AddProm",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PromController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PromController"],
		beego.ControllerComments{
			Method:           "UpdateProm",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PromController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:PromController"],
		beego.ControllerComments{
			Method:           "DeleteProm",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ReceiverController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ReceiverController"],
		beego.ControllerComments{
			Method:           "UpdateReceiver",
			Router:           `/:receiverid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ReceiverController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:ReceiverController"],
		beego.ControllerComments{
			Method:           "DeleteReceiver",
			Router:           `/:receiverid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:RuleController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:RuleController"],
		beego.ControllerComments{
			Method:           "SendAllRules",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:RuleController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:RuleController"],
		beego.ControllerComments{
			Method:           "AddRule",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:RuleController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:RuleController"],
		beego.ControllerComments{
			Method:           "UpdateRule",
			Router:           `/:ruleid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:RuleController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:RuleController"],
		beego.ControllerComments{
			Method:           "DeleteRule",
			Router:           `/:ruleid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAllUser",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:UserController"],
		beego.ControllerComments{
			Method:           "AddUser",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UpdatePassword",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers:UserController"],
		beego.ControllerComments{
			Method:           "DeleteUsers",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
