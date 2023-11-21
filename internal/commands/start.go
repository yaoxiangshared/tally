package commands

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"githup.com/tally/internal/config"
	"githup.com/tally/internal/core"
	"githup.com/tally/internal/event"
	"githup.com/tally/internal/server"
	"githup.com/tally/internal/service"
	"os"
	"os/signal"
	"syscall"
)

// StartCommand registers the start cli command.
var StartCommand = cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Starts web server",
	Action:  startAction,
}
var log = event.Log

// startAction start the web server and initializes the daemon
func startAction(ctx *cli.Context) error {
	//fmt.Printf("%#v\n", ctx)
	//c, err2 := ioutil.ReadFile("./cmd/web/load.yml")
	//fmt.Println(err2)
	//fmt.Println(string(c))
	a := "asSASA ddd dsjkdsjs dk"
	fmt.Println(len(a))
	b := "asSASA ddd dsjkdsjsこん dk"
	fmt.Println(len(b))
	fmt.Println(ctx.String("lang"))
	fmt.Println(ctx.Int("test"))
	fmt.Println("start web!")

	s := make([]int, 10)
	s1 := s[5:6]
	fmt.Println(len(s1))

	return nil
	conf := config.NewConfig(ctx)

	if conf.HttpPort() < 1 || conf.HttpPort() > 65535 {
		log.Fatal("server pojrt must be a number between 1 and 65535")
	}

	// pass this context down the chain
	//cctx, cancel := context.WithCancel(context.Background())

	if err := conf.Init(); err != nil {
		log.Fatal(err)
	}
	dbMysql := &core.DbMySql{}
	dbMysql.SetDb(conf.Db())
	core.SetDbProvider(dbMysql)
	service.SetConfig(conf)

	//userServant := service.UserServant{Db:conf.Db()}
	//user, _ := userServant.GetUser(1)
	//fmt.Println( (*user).Name)
	cctx, cancel := context.WithCancel(context.Background())
	go server.Start(cctx, conf)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Info("shutting down...")
	//conf.Shutdown()
	cancel()

	return nil
}
