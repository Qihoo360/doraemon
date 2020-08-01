package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

const baseControllers = "github.com/Qihoo360/doraemon/cmd/alert-gateway/controllers"

func init() {

	beego.GlobalControllerRouter[baseControllers+":AlertController"] = []beego.ControllerComments{
		{
			Method:           "GetAlerts",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "Confirm",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "HandleAlerts",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "ClassifyAlerts",
			Router:           `/classify`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "ShowAlerts",
			Router:           `/rules/:ruleid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

	beego.GlobalControllerRouter[baseControllers+":ConfigController"] = []beego.ControllerComments{
		{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "AddConfig",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "UpdateConfig",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "DeleteConfig",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

	beego.GlobalControllerRouter[baseControllers+":GroupController"] = []beego.ControllerComments{
		{

			Method:           "GetAllGroup",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "AddGroup",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "UpdateGroup",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "DeleteGroup",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

	beego.GlobalControllerRouter[baseControllers+":LoginController"] = []beego.ControllerComments{
		{
			Method:           "Local",
			Router:           `/local`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "Ldap",
			Router:           `/ldap`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "GetMethod",
			Router:           `/method`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "GetCurrentUser",
			Router:           `/username`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

	beego.GlobalControllerRouter[baseControllers+":LogoutController"] = []beego.ControllerComments{
		{
			Method:           "Logout",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

	beego.GlobalControllerRouter[baseControllers+":MaintainController"] = []beego.ControllerComments{
		{
			Method:           "GetAllMaintains",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "AddMaintain",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "UpdateMaintain",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "DeleteMaintain",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "GetHosts",
			Router:           `/:id/hosts`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

	beego.GlobalControllerRouter[baseControllers+":ManageController"] = []beego.ControllerComments{
		{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "AddManage",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "UpdateManage",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "DeleteManage",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

	beego.GlobalControllerRouter[baseControllers+":PlanController"] = []beego.ControllerComments{
		{
			Method:           "GetAllPlans",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "AddPlan",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "UpdatePlan",
			Router:           `/:planid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "DeletePlan",
			Router:           `/:planid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "GetAllReceiver",
			Router:           `/:planid/receivers/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "AddReceiver",
			Router:           `/:planid/receivers/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

	beego.GlobalControllerRouter[baseControllers+":PromController"] = []beego.ControllerComments{
		{
			Method:           "GetAllProms",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "AddProm",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "UpdateProm",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "DeleteProm",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

	beego.GlobalControllerRouter[baseControllers+":ReceiverController"] = []beego.ControllerComments{
		{
			Method:           "UpdateReceiver",
			Router:           `/:receiverid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "DeleteReceiver",
			Router:           `/:receiverid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

	beego.GlobalControllerRouter[baseControllers+":RuleController"] = []beego.ControllerComments{
		{
			Method:           "SendAllRules",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "AddRule",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "UpdateRule",
			Router:           `/:ruleid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "DeleteRule",
			Router:           `/:ruleid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

	beego.GlobalControllerRouter[baseControllers+":UserController"] = []beego.ControllerComments{
		{
			Method:           "GetAllUser",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "AddUser",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "UpdatePassword",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		{
			Method:           "DeleteUsers",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
	}

}
