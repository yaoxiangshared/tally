package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(router *gin.RouterGroup) {
	router.GET("/user-info", func(c *gin.Context) {

		c.JSON(http.StatusOK, 2)

	})
}
