package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	. "fGin/model"
	. "fGin/config"
	"github.com/gin-gonic/gin/json"
	. "fGin/ov"
	"strconv"
	"github.com/tidwall/gjson"
	"crypto/sha256"
	"encoding/hex"
)

func LoginAccount(c *gin.Context) {
	userName := c.Query("userName")
	password := c.Query("password")
	user := SysUser{Username: userName}
	Db.First(&user, user)
	m := sha256.New()
	m.Write([]byte(password)) // 需要加密的字符串为buf.String()
	pwdMd5 := hex.EncodeToString(m.Sum(nil))
	renderUser := User{
		Status:           "error",
		Type:             "account",
		CurrentAuthority: "",
	}
	if &user.Password != nil && user.Password == pwdMd5 {
		renderUser.Status = "ok"
		renderUser.CurrentAuthority = userName
		c.JSON(200, renderUser)
		return
	}
	c.JSON(200, renderUser)
}

func QueryMenu(c *gin.Context) {
	sysUser := SysUser{}
	sysUser.Username = c.Query("username")
	Db.First(&sysUser, sysUser)
	sysUserRole := []SysUserRole{}
	Db.Find(&sysUserRole, "user_id = ?", sysUser.UserId)
	roles := []int{}
	for _, r := range sysUserRole {
		roles = append(roles, r.RoleId)
	}
	menu := []SysMenu{}
	Db.Joins("left join sys_role_menu on sys_role_menu.menu_id=sys_menu.menu_id").Find(&menu, "parent_id = ? and type = ? and sys_role_menu.role_id in (?)", "0", "3", roles)
	menus := []Menu{}
	for _, m := range menu {
		mm := Menu{}
		mm.Name = m.Name
		mm.Icon = m.Icon
		mm.Path = m.Url
		mm.Authority = sysUser.Username
		menuChildren := []SysMenu{}
		menusChildren := []Menu{}
		Db.Find(&menuChildren, "parent_id = ?", m.MenuId)
		for _, mc := range menuChildren {
			mmm := Menu{}
			mmm.Name = mc.Name
			mmm.Icon = mc.Icon
			mmm.Path = mc.Url
			mmm.Authority = sysUser.Username
			menusChildren = append(menusChildren, mmm)
		}
		mm.Children = menusChildren
		menus = append(menus, mm)
	}
	render, _ := json.Marshal(menus)
	c.String(200, string(render))
}
func QueryTreeMenu(c *gin.Context) {
	//roleId, ex := c.GetQuery("roleId")
	//var allSysMenu = ""
	//sysUserMenu := []SysMenu{}
	//Db.Joins("left join sys_role_menu on sys_role_menu.menu_id = sys_menu.menu_id").Find(&sysUserMenu, "sys_role_menu.role_id = ? and sys_menu.type =?", roleId, "3")
	//sysUserMenujson, _ := json.Marshal(sysUserMenu)
	//allSysMenu += string(sysUserMenujson)
	//println(allSysMenu)
	treemenu := []Treemenu{}
	//获取所有菜单
	allMenu := []SysMenu{}
	Db.Find(&allMenu, "parent_id = ? and type = ?", "0", "3")
	//Db.Find(&allMenu, "parent_id = ? ", "0")
	for _, m := range allMenu {
		treemenu2 := []Treemenu{}
		t := Treemenu{}
		t.Title = m.Name
		t.Key = strconv.Itoa(m.MenuId)
		t.Selectable = true
		//if strings.Contains(allSysMenu, "\"MenuId\":"+strconv.Itoa(m.MenuId)) {
		//	t.Selectable = "true"
		//}
		menuChildren := []SysMenu{}
		Db.Find(&menuChildren, "parent_id = ?", m.MenuId)
		for _, mm := range menuChildren {
			t2 := Treemenu{}
			t2.Key = strconv.Itoa(mm.MenuId)
			t2.Title = mm.Name
			t.Selectable = true
			//if strings.Contains(allSysMenu, "\"MenuId\":"+strconv.Itoa(mm.MenuId)) {
			//	t2.Selectable = "true"
			//}
			treemenu2 = append(treemenu2, t2)
		}
		t.Children = treemenu2
		treemenu = append(treemenu, t)
	}
	c.JSON(200, treemenu)
}
func QuerySelectedKeys(c *gin.Context) {
	roleId, ex := c.GetQuery("roleId")
	if ex {
		allSysMenu := []string{}
		sysUserRoleMenu := []SysRoleMenu{}
		Db.Joins("left join sys_menu on sys_role_menu.menu_id = sys_menu.menu_id").Find(&sysUserRoleMenu, "sys_role_menu.role_id = ? and sys_menu.type!='0'  and sys_menu.type!='1'  and sys_menu.type!='2'", roleId)
		for _, m := range sysUserRoleMenu {
			allSysMenu = append(allSysMenu, strconv.Itoa(m.MenuId))
		}
		c.JSON(200, allSysMenu)
	}
}
func AddRole(c *gin.Context) {
	roleName := c.Query("RoleName")
	remark := c.Query("Remark")
	sysRole := SysRole{
		RoleName: roleName,
		Remark:   remark,
	}
	roleId := c.Query("RoleId")
	if roleId != "" {
		rid, _ := strconv.Atoi(roleId)
		sysRole.RoleId = rid
		Db.Where("role_id = ?", sysRole.RoleId).Model(&sysRole).Update(sysRole)
	} else {
		Db.Create(&sysRole)
	}

	c.String(200, "OK")
}
func UpdateRoleMenu(c *gin.Context) {
	roleId := c.Query("roleId")
	allSysMenu := []string{}
	sysUserRoleMenu := []SysRoleMenu{}
	Db.Find(&sysUserRoleMenu, "role_id = ?", roleId)
	for _, m := range sysUserRoleMenu {
		allSysMenu = append(allSysMenu, strconv.Itoa(m.MenuId))
	}
	c.JSON(200, allSysMenu)
}
func CurrentUser(c *gin.Context) {
	sysUser := SysUser{}
	sysUser.Username = c.Query("username")
	Db.First(&sysUser, sysUser)
	c.JSON(200, sysUser)
}
func QueryUser(c *gin.Context) {
	sysUser := []SysUser{}
	Db.Limit(10).Find(&sysUser)
	c.JSON(200, Page{List: sysUser, Pagination: Pagination{Total: 10}})
}
func QueryRole(c *gin.Context) {
	sysRole := []SysRole{}
	Db.Limit(10).Find(&sysRole)
	c.JSON(200, Page{List: sysRole, Pagination: Pagination{Total: 10}})
}
func UpDateRole(c *gin.Context) {
	sysRole := SysRole{}
	sysRole.RoleName = c.Query("roleName")
	sysRole.RoleId = c.GetInt("roleId")
	sysRole.Remark = c.Query("remark")
	Db.Where("role_id = ?", sysRole.RoleId).Model(&sysRole).Update(sysRole)
}
func DelRole(c *gin.Context) {
	ids := c.Query("roleIds")
	gjson.Parse(ids).ForEach(func(key, value gjson.Result) bool {
		id := value.Get("RoleId").String()
		Db.Unscoped().Delete(SysRoleMenu{}, "role_id = ?", id)
		Db.Unscoped().Delete(SysRole{}, "role_id = ?", id)
		return true
	})
}
func QueryByRoleId(c *gin.Context) {
	id := c.Query("roleId")
	sysRole := SysRole{}
	rid, _ := strconv.Atoi(id)
	sysRole.RoleId = rid
	Db.First(&sysRole, sysRole)
	c.JSON(200, sysRole)
}
