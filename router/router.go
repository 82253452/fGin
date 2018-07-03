package router

import (
	"gopkg.in/gin-gonic/gin.v1"
	. "fGin/config"
)

var Router *gin.Engine

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if E.HasPermissionForUser("admin", c.Request.RequestURI) == true {
			// permit alice to read data1
			println("1111111111111111")
		} else {
			// deny the request, show an error
			println("22222222222222222222")
		}
		c.Next()
	}
}
func init() {
	Router = gin.Default()
	Router.Static("/static", "./static/dist")
	Router.Use(AuthMiddleWare())
}
