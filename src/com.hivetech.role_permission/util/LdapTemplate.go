package util

import (
	"fmt"
	. "github.com/go-ldap/ldap/v3"
	"github.com/sirupsen/logrus"
	"pt_role_permission_go/src/com.hivetech.role_permission/config"
	"pt_role_permission_go/src/com.hivetech.role_permission/storage/model"
)

const (
	FIND_ALL = "(&(objectClass=User)(objectCategory=Person))"
)

var (
	FIND_ALL_BY_NAME = "(&(objectClass=User)(sAMAccountName=%s))"
	ATTRIBUTES       = []string{"dn", "sAMAccountName", "mail", "sn", "givenName"}
	ALL_ATT          = []string{}
	EMPTY_USER       = model.LdapUser{}
	EMPTY_LIST       []model.LdapUser
	QUERY_ERROR_MSG  = "Failed to query ldap: %w "
)

func connection() *Conn {
	myLdap, err := DialURL(config.LdapURL)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("ldap is close? ", myLdap.IsClosing())
	bindChecker := myLdap.Bind(config.LdapUserBind, config.LdapPassword)
	if bindChecker != nil {
		logrus.Error("Can't bind to user > ", config.LdapUsername)
		myLdap.Close()
		return nil
	}

	return myLdap
}

func FindAll(value string) []model.LdapUser {
	myLdap := connection()
	if myLdap == nil {
		return EMPTY_LIST
	}

	searchResult, errorSearch := myLdap.Search(searchRequest(FIND_ALL, value))
	defer myLdap.Close()
	if errorSearch != nil {
		logrus.Error(QUERY_ERROR_MSG, errorSearch)
		return EMPTY_LIST
	}

	return mapToUsers(searchResult.Entries)
}

func FindLdapUserByUsername(username string) model.LdapUser {
	myLdap := connection()
	if myLdap == nil {
		return EMPTY_USER
	}

	searchResult, errorSearch := myLdap.Search(searchRequest(FIND_ALL_BY_NAME, username))
	defer myLdap.Close()
	if errorSearch != nil {
		logrus.Error(QUERY_ERROR_MSG, errorSearch)
		return EMPTY_USER
	}

	if len(searchResult.Entries) == 0 {
		return EMPTY_USER
	}

	return toUser(searchResult.Entries[0])
}

func searchRequest(query string, params ...string) *SearchRequest {
	return NewSearchRequest(config.BaseDn, ScopeWholeSubtree, 0,
		0, 0, false,
		fmt.Sprintf(query, params), ALL_ATT, nil)
}

func toUser(entry *Entry) model.LdapUser {
	return model.LdapUser{
		Mail:           entry.GetAttributeValue("mail"),
		Name:           entry.GetAttributeValue("name"),
		Sn:             entry.GetAttributeValue("sn"),
		SAMAccountName: entry.GetAttributeValue("sAMAccountName"),
	}
}

func mapToUsers(entries []*Entry) (users []model.LdapUser) {
	for _, entry := range entries {
		users = append(users, toUser(entry))
	}

	return users
}
