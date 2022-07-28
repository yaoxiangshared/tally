package config

import "github.com/urfave/cli/v2"

var GlobalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "lang",
		Value: "english",
		Usage: "language for the greeting",
	},
	&cli.IntFlag{
		Name:  "pass",
		Value: 20,
		Usage: "password",
	},
	&cli.StringFlag{
		Name:  "mysql_host",
		Value: "aaa",
		Usage: "password",
	},
	&cli.IntFlag{
		Name:  "pass",
		Value: 20,
		Usage: "password",
	},
	&cli.IntFlag{
		Name:  "pass",
		Value: 20,
		Usage: "password",
	},
}
