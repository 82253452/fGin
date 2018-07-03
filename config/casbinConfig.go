package config

import (
	. "github.com/casbin/casbin"
	. "github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

var E *Enforcer

func init() {
	a := NewAdapter("mysql", "root:Fblife@20171019@tcp(47.92.100.148:3306)/fast4ward_dev", true)
	m := NewModel()
	m.AddDef("r", "r", "user, url")
	m.AddDef("p", "p", "role, url")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "r.url == p.url && myMatch(r.user,r.url)")
	E = NewEnforcer(m, a)
	E.AddFunction("myMatch", myMatch)
}
func myMatch(args ...interface{}) (interface{}, error) {
	println("匹配了")
	name1 := args[0].(string)
	name2 := args[1].(string)
	roles := E.GetRolesForUser(name1)
	for _, role := range roles {
		if E.HasPermissionForUser(role, name2) {
			return true, nil
		}
	}
	return false, nil
}
