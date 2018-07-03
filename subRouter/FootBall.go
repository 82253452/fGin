package subRouter

import (
	. "fGin/router"
	. "fGin/controller"
)

func init() {
	Router.GET("/footBall/saicheng/list", GetSaichengList)
	Router.GET("/footBall/jifen/list", GetJifenList)
}
