package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"githup.com/tally/internal/config"
	"githup.com/tally/internal/event"
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
	fmt.Println(ctx.String("lang"))
	fmt.Println(ctx.Int("test"))
	fmt.Println("start web!")
	conf := config.NewConfig(ctx)

	if conf.HttpPort() < 1 || conf.HttpPort() > 65535 {
		log.Fatal("server port must be a number between 1 and 65535")
	}

	// pass this context down the chain
	//cctx, cancel := context.WithCancel(context.Background())

	if err := conf.Init(); err != nil {
		log.Fatal(err)
	}
	//service.SetConfig(conf)
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
