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
		Value: "newskydb.mysql.rds.aliyuncs.com",
		Usage: "password",
	},
	&cli.StringFlag{
		Name:  "mysql_password",
		Value: "A60e316b7b85f1850",
		Usage: "password",
	},
	&cli.StringFlag{
		Name:  "mysql_user",
		Value: "tally_rw",
		Usage: "password",
	},
}
