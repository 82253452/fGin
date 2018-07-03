package subRouter

import (
	. "fGin/router"
	. "fGin/controller"
)

func init() {
	router := Router.Group("/sys")
	router.GET("/login", Login)
	router.GET("/sysInit", InitSysConfig)
	router.GET("/sysMenuTest", SysMeuTest)
	router.GET("/queryMenu", QueryMenu)
	router.GET("/currentUser", CurrentUser)
	router.GET("/login/account", LoginAccount)
	router.GET("/queryUser", QueryUser)
	router.GET("/queryRole", QueryRole)
	router.GET("/queryTreeMenu", QueryTreeMenu)
	router.GET("/querySelectedKeys", QuerySelectedKeys)
	router.GET("/addRole", AddRole)
	router.GET("/upDateRole", UpDateRole)
	router.GET("/delRole", DelRole)
	router.GET("/updateRoleMenu", UpdateRoleMenu)
	router.GET("/queryByRoleId", QueryByRoleId)
}
