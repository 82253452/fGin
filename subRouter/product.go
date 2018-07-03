package subRouter

import (
	. "fGin/router"
	. "fGin/controller"
)

func init() {
	Router.GET("/", Test)
}
