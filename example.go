package main

import (
	"fmt"
	"log"

	//"fmt"
	//"github.com/gin-gonic/gin"
	//"github.com/urfave/cli"
	//"log"
	//"net/http"
	"os"
	//"strings"
	"github.com/urfave/cli/v2"
)

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hel≤lo World"))
	//})
	//
	//http.Serve("127.0.0.1:8000", nil)
	//r := gin.Default()
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	//截取/
	//	action = strings.Trim(action, "/")
	//	c.String(http.StatusOK, name+" is "+action)
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	app := &cli.App{
		Name: "greet",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
