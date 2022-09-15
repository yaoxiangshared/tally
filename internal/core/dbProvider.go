package core

import (
	"gorm.io/gorm"
	"log"
)

type DbProvider interface {
	Db() *gorm.DB
}

var dbProvider DbProvider

// SetDbProvider sets the provider to get a gorm db connection.
func SetDbProvider(provider DbProvider) {
	dbProvider = provider
}

func GetDbProvider() DbProvider {
	return dbProvider
}

//数据库连接的全局变量

type DbMySql struct {
	db *gorm.DB
}

// Db returns the db connection.
func (c *DbMySql) Db() *gorm.DB {
	if c.db == nil {
		log.Fatal("config: database not connected")
	}

	return c.db
}

func (c *DbMySql) SetDb(db *gorm.DB) {
	c.db = db
}
