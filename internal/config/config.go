package config

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"githup.com/tally/internal/event"
	"sync"
)

var once sync.Once
var log = event.Log

type Config struct {
	once     sync.Once
	db       *gorm.DB
	options  *Options
	//settings *Settings
	token    string
	serial   string
}
// NewConfig initialises a new configuration file
func NewConfig(ctx *cli.Context) *Config {
	initLogger(ctx.Bool("debug"))

	c := &Config{
		options: NewOptions(ctx),
		//token:   rnd.Token(8),
	}

	return c
}

func initLogger(debug bool) {
	once.Do(func() {
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})

		if debug {
			log.SetLevel(logrus.DebugLevel)
		} else {
			log.SetLevel(logrus.InfoLevel)
		}
	})
}