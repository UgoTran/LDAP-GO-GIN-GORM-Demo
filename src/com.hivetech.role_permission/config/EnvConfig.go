package config

import (
	"fmt"
	. "github.com/go-ldap/ldap/v3"
	"github.com/spf13/viper"
	"strings"
)

var (
	LdapConnection *Conn
	LdapURL        = ""
	BaseDn         = ""
	LdapUsername   = ""

	LdapUserBind = ""
	LdapPassword = ""
)

func InitEnvConfig() {
	LdapURL = fmt.Sprintf("%v", viper.Get("LDAP_URL"))
	BaseDn = fmt.Sprintf("%v", viper.Get("LDAP_BASE_DN"))
	LdapUsername = fmt.Sprintf("%v", viper.Get("LDAP_UN"))
	LdapUserBind = strings.Join([]string{LdapUsername, BaseDn}, ",")
	LdapPassword = fmt.Sprintf("%v", viper.Get("LDAP_PW"))
}
