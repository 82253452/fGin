package model

import (
	"github.com/jinzhu/gorm"
	. "fGin/config"
)

/**
/球队
 */
type FootBallTeam struct {
	gorm.Model
	Name       string
	Img        string
	GameNumber uint   //赛
	Win        uint   //赢
	Lost       uint   //输
	Goal       uint   //进球
	Fumble     uint   //失球
	Integral   uint   //积分
	Draw       uint   //平局
	Group      string //组别
}

func init() {
	Db.AutoMigrate(&FootBallTeam{})
}
