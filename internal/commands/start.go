package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"githup.com/tally/internal/config"
	"githup.com/tally/internal/event"
	"githup.com/tally/internal/service"
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
	fmt.Println(ctx.String("lang"))
	fmt.Println(ctx.Int("pass"))
	fmt.Println("start web!")
	conf := config.NewConfig(ctx)
	service.SetConfig(conf)
	//cctx, cancel := context.WithCancel(context.Background())
	//go server.Start(cctx, conf)
	//
	//quit := make(chan os.Signal)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//
	//<-quit
	//log.Info("shutting down...")
	////conf.Shutdown()
	//cancel()

	return nil
}
