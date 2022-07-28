package main

import (
	"github.com/urfave/cli/v2"
	"githup.com/tally/internal/commands"
	"githup.com/tally/internal/config"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "develop"
	app.Flags = config.GlobalFlags
	//app.Flags=[]cli.Flag{
	//	&cli.StringFlag{
	//		Name: "lang",
	//		Value: "english",
	//		Usage: "language for the greeting",
	//	},
	//	&cli.IntFlag{
	//		Name: "pass",
	//		Value: 20,
	//		Usage: "password",
	//	},
	//}
	app.Commands = []*cli.Command{
		&commands.StartCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
