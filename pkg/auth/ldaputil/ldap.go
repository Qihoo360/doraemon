package ldaputil

import (
	"fmt"
	"github.com/Qihoo360/doraemon/cmd/alert-gateway/logs"
	"github.com/go-ldap/ldap"
	"github.com/pkg/errors"
)

//[auth.ldap]
//enabled = false
//ldap_url = ldap://127.0.0.1:3893
//ldap_search_dn = "hackers,ou=superheros,dc=glauth,dc=com"
//ldap_search_password = dogood
//ldap_base_dn = "dc=glauth,dc=com"
//ldap_filter = "(&(ou=superheros)(mail=%s))"
//ldap_uid = cn
//ldap_scope = 2
//ldap_connection_timeout = 30

var (
	DefaultClient   *LdapClient
	ErrLdapAuthFail = errors.New("user does not exist or too many entries returned")
	ErrLdapUnInit   = errors.New("ldap is not initialized")
)

type LdapConfig struct {
	Url          string
	BindUsername string
	BindPassword string
	Scope        int
	BaseDN       string
	Filter       string
}

type LdapClient struct {
	*LdapConfig
}

func InitLdap(c *LdapConfig) {
	if DefaultClient == nil {
		DefaultClient = &LdapClient{
			LdapConfig: c,
		}
	}
}

func Authenticate(username, password string) error {
	if DefaultClient == nil {
		return ErrLdapUnInit
	}

	return DefaultClient.Authenticate(username, password)
}

func (c *LdapClient) Authenticate(username, password string) error {

	conn, err := ldap.DialURL(c.Url)
	if err != nil {
		return err
	}
	defer conn.Close()

	// First bind with a read only user
	if err = conn.Bind(c.BindUsername, c.BindPassword); err != nil {
		return errors.Wrap(err, "first bind error")
	}

	filter := fmt.Sprintf(c.Filter, username)

	// Search for the given username
	req := ldap.NewSearchRequest(
		c.BaseDN,
		c.Scope,
		ldap.NeverDerefAliases,
		0, 0, false,
		filter,
		[]string{"dn"},
		nil,
	)

	sr, err := conn.Search(req)
	if err != nil {
		return errors.Wrap(err, "search error")
	}

	if len(sr.Entries) != 1 {
		logs.Debug("len(sr.Entries)=%d", len(sr.Entries))
		return ErrLdapAuthFail
	}

	userDN := sr.Entries[0].DN
	if err = conn.Bind(userDN, password); err != nil {
		return errors.Wrap(err, "second bind error")
	}

	return nil
}
