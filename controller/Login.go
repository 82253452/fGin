package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	. "fGin/config"
	. "fGin/model"
	"encoding/hex"
	"crypto/sha256"
	"strconv"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	pwd := c.Query("pwd")
	user := SysUser{Username: username}
	Db.First(&user, user)
	m := sha256.New()
	m.Write([]byte(pwd)) // 需要加密的字符串为buf.String()
	pwdMd5 := hex.EncodeToString(m.Sum(nil))
	if &user.Password != nil && user.Password == pwdMd5 {
		println("yes")
		sysRole := []SysRole{}
		Db.Joins("left join sys_user_role on sys_role.role_id=sys_user_role.role_id").Where("user_id = ?", user.UserId).Find(&sysRole)
		for _, role := range sysRole {
			E.AddRoleForUser(username, strconv.Itoa(role.RoleId))
			sysMenu := [] SysMenu{}
			Db.Joins("left join sys_role_menu on sys_role_menu.menu_id = sys_menu.menu_id").Where("sys_role_menu.role_id = ?", strconv.Itoa(role.RoleId)).Find(&sysMenu)
			for _, menu := range sysMenu {
				E.AddPermissionForUser(username, menu.Url)
			}
		}
	}
}

func InitSysConfig(c *gin.Context) {
	sysUserRole := []SysUserRole{}
	Db.Find(&sysUserRole)
	for _, userRole := range sysUserRole {
		E.AddRoleForUser(strconv.Itoa(userRole.UserId), strconv.Itoa(userRole.RoleId))
	}

	roles := []SysRole{}
	Db.Find(&roles)
	for _, role := range roles {
		sysRoleMenu := []SysRoleMenu{}
		Db.Find(&sysRoleMenu, "role_id = ?", role.RoleId)
		for _, roleMenu := range sysRoleMenu {
			sysMenu := SysMenu{}
			sysMenu.MenuId = roleMenu.MenuId
			Db.First(&sysMenu, sysMenu)
			if &sysMenu.Url != nil && sysMenu.Url != "" {
				E.AddPolicy(strconv.Itoa(role.RoleId), sysMenu.Url)
			}
		}
	}
	c.String(200, "hahahahahahahanice")
}
func SysMeuTest(c *gin.Context) {
	println(E.HasRoleForUser("2", "1"))//用户是否有角色
	println(E.HasPermissionForUser("1", "sys/user.html")) //角色是否有权限
	println(E.HasGroupingPolicy("2", "1"))//用户是否有角色
	println(E.Enforce("3", "sys/user.html"))//用户是否有权限
	c.String(200,"234242")
}
