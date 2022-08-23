package config

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

var GlobalFlags = []cli.Flag{
	altsrc.NewStringFlag(&cli.StringFlag{
		Name:  "lang",
		Value: "english",
		Usage: "language for the greeting",
	}),

	altsrc.NewStringFlag(&cli.StringFlag{
		Name:  "config_path",
		Value: "",
		Usage: "配置文件",
	}),
	//altsrc.NewIntFlag(&cli.IntFlag{
	//	Name:  "pass",
	//	Value: 20,
	//	Usage: "password",
	//}),
	//altsrc.NewStringFlag(&cli.StringFlag{
	//	Name:  "mysql_host",
	//	Value: "aaaa",
	//	Usage: "password",
	//}),
	//altsrc.NewStringFlag(&cli.StringFlag{
	//	Name:  "mysql_password",
	//	Value: "bbb",
	//	Usage: "password",
	//}),
	//altsrc.NewStringFlag(&cli.StringFlag{
	//	Name:  "mysql_user",
	//	Value: "aaa",
	//	Usage: "password",
	//}),
	&cli.StringFlag{
		Name: "load",
	},
	&cli.StringFlag{
		Name: "setting",
	},
	altsrc.NewIntFlag(&cli.IntFlag{
		Name:  "test",
		Value: 1,
	}),
}
