package server

import (
	"github.com/gin-gonic/gin"
	"githup.com/tally/internal/api"
	"githup.com/tally/internal/config"
)

func registerRoutes(router *gin.Engine, conf *config.Config) {

	user := router.Group("user")
	{
		// Config options.
		api.GetUserInfo(user)

	}

}
