package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"githup.com/tally/internal/event"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once
var log = event.Log

type Config struct {
	once    sync.Once
	db      *gorm.DB
	options *Options
	//settings *Settings
	token  string
	serial string
}

// NewConfig initialises a new configuration file
func NewConfig(ctx *cli.Context) *Config {
	initLogger(ctx.Bool("debug"))

	c := &Config{
		options: NewOptions(ctx),
		//token:   rnd.Token(8),
	}
	var err error
	configFilePath, err := configFile(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configFile)
	if err = c.options.Load(configFilePath); err != nil {
		log.Warnf("config: %s", err)
	} else {
		log.Debugf("config: options loaded from %s", err)
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

func configFile(ctx *cli.Context) (string, error) {
	path := ctx.String("config_path")
	if path == "" {
		return path, fmt.Errorf("config path is empty")
	}
	return path, nil

}

// Init creates directories, parses additional config files, opens a database connection and initializes dependencies.
func (c *Config) Init() error {
	//if err := c.CreateDirectories(); err != nil {
	//	return err
	//}
	//
	//if err := c.initStorage(); err != nil {
	//	return err
	//}
	//
	//if insensitive, err := c.CaseInsensitive(); err != nil {
	//	return err
	//} else if insensitive {
	//	log.Infof("config: case-insensitive file system detected")
	//	fs.IgnoreCase()
	//}
	//
	//if cpuName := cpuid.CPU.BrandName; cpuName != "" {
	//	log.Debugf("config: running on %s, %s memory detected", txt.Quote(cpuid.CPU.BrandName), humanize.Bytes(TotalMem))
	//}
	//
	//// Check memory requirements.
	//if TotalMem < 128*Megabyte {
	//	return fmt.Errorf("config: %s of memory detected, %d GB required", humanize.Bytes(TotalMem), MinMem/Gigabyte)
	//} else if LowMem {
	//	log.Warnf(`config: less than %d GB of memory detected, please upgrade if server becomes unstable or unresponsive`, MinMem/Gigabyte)
	//}
	//
	//// Show swap info.
	//if TotalMem < RecommendedMem {
	//	log.Infof("config: make sure your server has enough swap configured to prevent restarts when there are memory usage spikes")
	//}
	//
	//c.initSettings()
	//c.initHub()
	//
	//c.Propagate()

	return c.connectDb()
}
