package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"githup.com/tally/internal/commands"
	"githup.com/tally/internal/config"
	"log"
	"os"
)

func main() {

	type Rope string
	var a Rope = "aaa"
	fmt.Printf("Rope is %s\n", a)
	app := cli.NewApp()
	app.Version = "develop"
	app.Flags = config.GlobalFlags

	app.Before = altsrc.InitInputSourceWithContext(config.GlobalFlags, altsrc.NewYamlSourceFromFlagFunc("setting"))
	//app.Flags = config.GlobalFlags
	app.Commands = []*cli.Command{
		&commands.StartCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
