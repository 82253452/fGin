package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"fGin/model"
	. "fGin/config"
)

func Test(c *gin.Context) {
	Db.Save(model.Product{Code: "123", Price: 123})
	c.String(http.StatusOK, "Hello word")
}
