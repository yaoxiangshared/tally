package main

import (
	"github.com/urfave/cli/v2"
	"githup.com/tally/internal/commands"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Flags=[]cli.Flag{
		&cli.StringFlag{
			Name: "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	app.Commands = []*cli.Command{
		&commands.StartCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
