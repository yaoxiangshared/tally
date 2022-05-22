package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetData(router *gin.RouterGroup) {
	router.GET("/dashboard", func(c *gin.Context) {

		c.JSON(http.StatusOK, 1)

	})
}
