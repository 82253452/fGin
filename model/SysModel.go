package model

import "time"

type SysUser struct {
	UserId         int
	Uid            int
	Username       string
	Password       string
	Email          string
	Mobile         string
	Unionid        string
	Avatar         string
	Status         int
	Create_user_id int
	Create_time    time.Time
	Last_login     int
}
type SysRole struct {
	RoleId       int
	RoleName     string
	Remark       string
	CreateUserId int
	CreateTime   time.Time
}
type SysUserRole struct {
	Id     int
	UserId int
	RoleId int
}
type SysMenu struct {
	MenuId   int
	ParentId int
	Name     string
	Url      string
	Perms    string
	MenuType int  `gorm:"column:type"`
	Icon     string
	OrderNum int
}
type SysRoleMenu struct {
	Id     int
	RoleId int
	MenuId int
}
type User struct {
	Status           string `json:"status"`
	Type             string `json:"type"`
	CurrentAuthority string `json:"currentAuthority"`
}
