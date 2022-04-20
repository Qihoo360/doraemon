package main

import (
	"doraemon/cmd/alert-gateway/initial"
	"doraemon/pkg/auth/ldaputil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-ldap/ldap"

	_ "doraemon/cmd/alert-gateway/logs"
	_ "doraemon/cmd/alert-gateway/routers"
)

func parseLdapScope(scope string) int {
	switch scope {
	case "0":
		return ldap.ScopeBaseObject
	case "1":
		return ldap.ScopeSingleLevel
	case "2":
		return ldap.ScopeWholeSubtree
	default:
		return ldap.ScopeWholeSubtree
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	cfg, err := beego.AppConfig.GetSection("auth.ldap")
	if err != nil {
		logs.Warn("ldap config error: %v, ldap will not be supported", err)
	} else if cfg != nil && cfg["enabled"] == "true" {
		ldapCfg := ldaputil.LdapConfig{
			Url:          cfg["ldap_url"],
			BaseDN:       cfg["ldap_base_dn"],
			Scope:        parseLdapScope(cfg["ldap_scope"]),
			BindUsername: cfg["ldap_search_dn"],
			BindPassword: cfg["ldap_search_password"],
			Filter:       cfg["ldap_filter"],
		}
		ldaputil.InitLdap(&ldapCfg)
	}

	initial.InitDb()
	beego.Run()
}

//go:generate sh -c "echo 'package routers; import \"github.com/astaxie/beego\"; func init() {beego.BConfig.RunMode = beego.DEV}' > routers/0.go"
//go:generate sh -c "echo 'package routers; import \"os\"; func init() {os.Exit(0)}' > routers/z.go"
//go:generate go run $GOFILE
//go:generate sh -c "rm routers/0.go routers/z.go"
