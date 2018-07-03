package main

import (
	. "github.com/casbin/casbin"
	. "github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a := NewAdapter("mysql", "root:Fblife@20171019@tcp(47.92.100.148:3306)/abc")
	e := NewEnforcer("C:\\Users\\admin\\go\\src\\github.com\\casbin\\casbin\\examples\\rbac_model.conf", a)
	e.LoadPolicy()
	// Check the permission.
	sub := "alice"
	obj := "data1"
	act := "read"

	if e.Enforce(sub, obj, act) == true {
		// permit alice to read data1
		println("1111111111111111")
	} else {
		// deny the request, show an error
		println("22222222222222222222")
	}

	e.SavePolicy()
}
