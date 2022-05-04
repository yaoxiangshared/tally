package service

import (


	gc "github.com/patrickmn/go-cache"
	"githup.com/tally/internal/config"
)

var conf *config.Config

var services struct {
	FolderCache *gc.Cache
	CoverCache  *gc.Cache
	ThumbCache  *gc.Cache
}

func SetConfig(c *config.Config) {
	if c == nil {
		panic("config is nil")
	}

	conf = c

	//photoprism.SetConfig(c)
}

func Config() *config.Config {
	if conf == nil {
		panic("config is nil")
	}

	return conf
}
