package common

import (
	"crypto/tls"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	goldap "gopkg.in/ldap.v2"
)

type LdapUserInfo struct {
	Name    string
	Email   string
	Display string
}

type LdapConf struct {
	LdapURL               string `json:"ldap_url"`
	LdapSearchDn          string `json:"ldap_search_dn"`
	LdapSearchPassword    string `json:"ldap_search_password"`
	LdapBaseDn            string `json:"ldap_base_dn"`
	LdapFilter            string `json:"ldap_filter"`
	LdapUID               string `json:"ldap_uid"`
	LdapScope             int    `json:"ldap_scope"`
	LdapConnectionTimeout int    `json:"ldap_connection_timeout"`
}

// LdapUser ...
type LdapUser struct {
	Username string `json:"ldap_username"`
	Email    string `json:"ldap_email"`
	Realname string `json:"ldap_realname"`
	DN       string `json:"-"`
}

// ValidateLdapConf ...
func ValidateLdapConf(ldapConfs LdapConf) (LdapConf, error) {
	var err error

	if ldapConfs.LdapURL == "" {
		return ldapConfs, fmt.Errorf("can not get any available LDAP_URL")
	}

	ldapConfs.LdapURL, err = formatLdapURL(ldapConfs.LdapURL)

	if err != nil {
		//logs.Error("invalid LdapURL format, error: %v", err)
		return ldapConfs, err
	}

	// Compatible with legacy codes
	// in previous harbor.cfg:
	// the scope to search for users, 1-LDAP_SCOPE_BASE, 2-LDAP_SCOPE_ONELEVEL, 3-LDAP_SCOPE_SUBTREE
	switch ldapConfs.LdapScope {
	case 1:
		ldapConfs.LdapScope = goldap.ScopeBaseObject
	case 2:
		ldapConfs.LdapScope = goldap.ScopeSingleLevel
	case 3:
		ldapConfs.LdapScope = goldap.ScopeWholeSubtree
	default:
		return ldapConfs, fmt.Errorf("invalid ldap search scope")
	}

	//	value := reflect.ValueOf(ldapConfs)
	//	lType := reflect.TypeOf(ldapConfs)
	//	for i := 0; i < value.NumField(); i++ {
	//		fmt.Printf("Field %d: %v %v\n", i, value.Field(i), lType.Field(i).Name)
	//	}

	return ldapConfs, nil

}

func formatLdapURL(ldapURL string) (string, error) {

	var protocol, hostport string
	var err error

	_, err = url.Parse(ldapURL)
	if err != nil {
		return "", fmt.Errorf("parse Ldap Host ERR: %s", err)
	}

	if strings.Contains(ldapURL, "://") {
		splitLdapURL := strings.Split(ldapURL, "://")
		protocol, hostport = splitLdapURL[0], splitLdapURL[1]
		if !((protocol == "ldap") || (protocol == "ldaps")) {
			return "", fmt.Errorf("unknown ldap protocol")
		}
	} else {
		hostport = ldapURL
		protocol = "ldap"
	}

	if strings.Contains(hostport, ":") {
		splitHostPort := strings.Split(hostport, ":")
		port, error := strconv.Atoi(splitHostPort[1])
		if error != nil {
			return "", fmt.Errorf("illegal url port")
		}
		if port == 636 {
			protocol = "ldaps"
		}

	} else {
		switch protocol {
		case "ldap":
			hostport = hostport + ":389"
		case "ldaps":
			hostport = hostport + ":636"
		}
	}

	fLdapURL := protocol + "://" + hostport

	return fLdapURL, nil

}

// Bind establish a connection to ldap based on ldapConfs and bind the user with given parameters.
func Bind(ldapConfs LdapConf, dn string, password string) error {
	conn, err := dialLDAP(ldapConfs)
	if err != nil {
		return err
	}
	defer conn.Close()
	if ldapConfs.LdapSearchDn != "" {
		if err := bindLDAPSearchDN(ldapConfs, conn); err != nil {
			return err
		}
	}
	return conn.Bind(dn, password)
}

// MakeFilter ...
func MakeFilter(username string, ldapFilter string, ldapUID string) string {

	var filterTag string

	if username == "" {
		filterTag = "*"
	} else {
		filterTag = username
	}

	if ldapFilter == "" {
		ldapFilter = "(" + ldapUID + "=" + filterTag + ")"
	} else {
		if !strings.Contains(ldapFilter, ldapUID+"=") {
			ldapFilter = "(&" + ldapFilter + "(" + ldapUID + "=" + filterTag + "))"
		} else {
			ldapFilter = strings.Replace(ldapFilter, ldapUID+"=*", ldapUID+"="+filterTag, -1)
		}
	}

	return ldapFilter
}

// SearchUser ...
func SearchUser(ldapConfs LdapConf) ([]LdapUser, error) {
	var ldapUsers []LdapUser
	var ldapConn *goldap.Conn
	var err error

	ldapConn, err = dialLDAP(ldapConfs)

	if err != nil {
		return nil, err
	}
	defer ldapConn.Close()

	if ldapConfs.LdapSearchDn != "" {
		err = bindLDAPSearchDN(ldapConfs, ldapConn)
		if err != nil {
			return nil, err
		}
	}

	if ldapConfs.LdapBaseDn == "" {
		return nil, fmt.Errorf("can not get any available LDAP_BASE_DN")
	}

	result, err := searchLDAP(ldapConfs, ldapConn)

	if err != nil {
		return nil, err
	}

	for _, ldapEntry := range result.Entries {
		var u LdapUser
		for _, attr := range ldapEntry.Attributes {
			val := attr.Values[0]
			//logs.Info("Current ldap entry attr name: %s\n", attr.Name)
			switch strings.ToLower(attr.Name) {
			case strings.ToLower(ldapConfs.LdapUID):
				u.Username = val
			case "uid":
				u.Realname = val
			case "cn":
				u.Realname = val
			case "mail":
				u.Email = val
			case "email":
				u.Email = val
			}
		}
		u.DN = ldapEntry.DN
		ldapUsers = append(ldapUsers, u)
	}

	return ldapUsers, nil
}

func dialLDAP(ldapConfs LdapConf) (*goldap.Conn, error) {

	var err error
	var ldap *goldap.Conn
	splitLdapURL := strings.Split(ldapConfs.LdapURL, "://")
	protocol, hostport := splitLdapURL[0], splitLdapURL[1]

	// Sets a Dial Timeout for LDAP
	connectionTimeout := ldapConfs.LdapConnectionTimeout
	goldap.DefaultTimeout = time.Duration(connectionTimeout) * time.Second

	switch protocol {
	case "ldap":
		ldap, err = goldap.Dial("tcp", hostport)
	case "ldaps":
		ldap, err = goldap.DialTLS("tcp", hostport, &tls.Config{InsecureSkipVerify: true})
	}

	return ldap, err
}

func bindLDAPSearchDN(ldapConfs LdapConf, ldap *goldap.Conn) error {

	var err error

	ldapSearchDn := ldapConfs.LdapSearchDn
	ldapSearchPassword := ldapConfs.LdapSearchPassword

	err = ldap.Bind(ldapSearchDn, ldapSearchPassword)
	if err != nil {
		//logs.Info("Bind search dn error", err)
		return err
	}

	return nil
}

func searchLDAP(ldapConfs LdapConf, ldap *goldap.Conn) (*goldap.SearchResult, error) {

	var err error
	ldapBaseDn := ldapConfs.LdapBaseDn
	ldapScope := ldapConfs.LdapScope
	ldapFilter := ldapConfs.LdapFilter

	attributes := []string{"uid", "cn", "mail", "email"}
	lowerUID := strings.ToLower(ldapConfs.LdapUID)
	if lowerUID != "uid" && lowerUID != "cn" && lowerUID != "mail" && lowerUID != "email" {
		attributes = append(attributes, ldapConfs.LdapUID)
	}
	searchRequest := goldap.NewSearchRequest(
		ldapBaseDn,
		ldapScope,
		goldap.NeverDerefAliases,
		0,     // Unlimited results.
		0,     // Search Timeout.
		false, // Types Only
		ldapFilter,
		attributes,
		nil,
	)

	result, err := ldap.Search(searchRequest)

	if err != nil {
		//logs.Info("LDAP search error", err)
		return nil, err
	}

	return result, nil
}

func Authenticate(m AuthModel) (*LdapUserInfo, error) {
	u := m.Username
	if len(strings.TrimSpace(u)) == 0 {
		//logs.Info("LDAP authentication failed for empty user id.")
		return nil, fmt.Errorf("LDAP authentication failed for empty user id.")
	}
	const metaChars = "&|!=~*<>()"

	for _, c := range metaChars {
		if strings.ContainsRune(u, c) {
			return nil, fmt.Errorf("the principal contains meta char: %q", c)
		}
	}

	ldapConfs := LdapConf{}

	section, err := beego.AppConfig.GetSection("auth.ldap")
	if err != nil {
		return nil, fmt.Errorf("Can't find ldap config. ")
	}

	ldapConfs.LdapURL = section["ldap_url"]
	ldapConfs.LdapSearchDn = section["ldap_search_dn"]
	ldapConfs.LdapSearchPassword = section["ldap_search_password"]
	ldapConfs.LdapBaseDn = section["ldap_base_dn"]
	ldapConfs.LdapFilter = section["ldap_filter"]
	ldapConfs.LdapUID = section["ldap_uid"]
	ldapScope, err := strconv.ParseInt(section["ldap_scope"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("ldap_scope parse error, must be int. ")
	}
	ldapConnectionTimeout, err := strconv.ParseInt(section["ldap_connection_timeout"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("ldap_connection_timeout parse error, must be int. ")
	}
	ldapConfs.LdapScope = int(ldapScope)
	ldapConfs.LdapConnectionTimeout = int(ldapConnectionTimeout)
	ldapConfs, err = ValidateLdapConf(ldapConfs)

	if err != nil {
		return nil, fmt.Errorf("invalid ldap request: %v", err)
	}

	ldapConfs.LdapFilter = MakeFilter(u, ldapConfs.LdapFilter, ldapConfs.LdapUID)
	ldapUsers, err := SearchUser(ldapConfs)

	if err != nil {
		//logs.Warning("ldap search fail: %v", err)
		return nil, err
	}

	if len(ldapUsers) == 0 {
		//logs.Warning("Not found an entry.")
		return nil, fmt.Errorf("Not found an entry. ")
	} else if len(ldapUsers) != 1 {
		//logs.Warning("Found more than one entry.")
		return nil, fmt.Errorf("Found more than one entry. ")
	}

	user := &LdapUserInfo{
		Name:    ldapUsers[0].Username,
		Email:   ldapUsers[0].Email,
		Display: ldapUsers[0].Realname,
	}

	dn := ldapUsers[0].DN

	//logs.Info("username: %s, dn: %s", user.Name, dn)
	if err := Bind(ldapConfs, dn, m.Password); err != nil {
		//logs.Warning("Failed to bind user, username: %s, dn: %s, error: %v", user.Name, dn, err)
		return nil, fmt.Errorf("Failed to bind user, username: %s, dn: %s, error: %v ", user.Name, dn, err)
	}

	return user, nil
}
