package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "develop"
	//app.Flags = config.GlobalFlags
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

	flags := []cli.Flag{
		altsrc.NewIntFlag(&cli.IntFlag{Name: "test"}),
		&cli.StringFlag{Name: "load"},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("yaml ist rad")
		fmt.Println(c.Int("test"))
		fmt.Println(c.String("load"))
		return nil
	}

	app.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("load"))
	app.Flags = flags
	//app.Commands = []*cli.Command{
	//	&commands.StartCommand,
	//}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
