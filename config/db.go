package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() () {
	Db, _ = gorm.Open("mysql", "root:Fblife@20171019@tcp(47.92.100.148:3306)/fast4ward_dev?charset=utf8&parseTime=True")
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.LogMode(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		rs := []rune(defaultTableName)
		return string(rs[0:len(rs)-1])
	}
}
func GetDb() *gorm.DB {
	Lab: err := Db.DB().Ping()
	if err != nil {
		goto Lab
	}
	return Db
}
