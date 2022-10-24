package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"githup.com/tally/internal/core"
	"githup.com/tally/internal/service"
	"net/http"
)

func GetUserInfo(router *gin.RouterGroup) {
	router.GET("/user-info", func(c *gin.Context) {
		fmt.Println(222)
		userServant := service.UserServant{Db: core.GetDbProvider().Db()}
		user, _ := userServant.GetUser(1)
		fmt.Printf("%#v\n", user)
		fmt.Println((user).Name)
		fmt.Println(111)
		c.JSON(http.StatusOK, user)

	})
}
