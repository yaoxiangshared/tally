package server

import (
	"github.com/gin-gonic/gin"
	"githup.com/tally/internal/api"
	"githup.com/tally/internal/config"
)

func registerRoutes(router *gin.Engine, conf *config.Config) {

	v1 := router.Group("")
	{
		// Config options.
		api.GetData(v1)
		api.GetUserInfo(v1)

	}
}
