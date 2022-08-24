package config

import (
	"errors"
	"fmt"
	"githup.com/tally/internal/mutex"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	//"github.com/photoprism/photoprism/internal/entity"
	//"github.com/photoprism/photoprism/internal/mutex"
)

var dsnPattern = regexp.MustCompile(
	`^(?:(?P<user>.*?)(?::(?P<password>.*))?@)?` +
		`(?:(?P<net>[^\(]*)(?:\((?P<server>[^\)]*)\))?)?` +
		`\/(?P<name>.*?)` +
		`(?:\?(?P<params>[^\?]*))?$`)

// DatabaseDriver returns the database driver name.
func (c *Config) DatabaseDriver() string {
	switch strings.ToLower(c.options.DatabaseDriver) {
	case MySQL, MariaDB:
		c.options.DatabaseDriver = MySQL
	case SQLite, "sqlite", "sqllite", "test", "file", "":
		c.options.DatabaseDriver = SQLite
	case "tidb":
		log.Warnf("config: database driver 'tidb' is deprecated, using sqlite")
		c.options.DatabaseDriver = SQLite
		c.options.DatabaseDsn = ""
	default:
		log.Warnf("config: unsupported database driver %s, using sqlite", c.options.DatabaseDriver)
		c.options.DatabaseDriver = SQLite
		c.options.DatabaseDsn = ""
	}

	return c.options.DatabaseDriver
}

// DatabaseDsn returns the database data source name (DSN).
func (c *Config) DatabaseDsn() string {
	if c.options.DatabaseDsn == "" {
		switch c.DatabaseDriver() {
		case MySQL, MariaDB:
			return fmt.Sprintf(
				"%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&collation=utf8mb4_unicode_ci&parseTime=true",
				c.DatabaseUser(),
				c.DatabasePassword(),
				c.DatabaseServer(),
				c.DatabaseName(),
			)
		//case Postgres:
		//	return fmt.Sprintf(
		//		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=UTC",
		//		c.DatabaseUser(),
		//		c.DatabasePassword(),
		//		c.DatabaseName(),
		//		c.DatabaseHost(),
		//		c.DatabasePort(),
		//	)
		//case SQLite:
		//	return filepath.Join(c.StoragePath(), "index.db")
		default:
			log.Errorf("config: empty database dsn")
			return ""
		}
	}

	return c.options.DatabaseDsn
}

// ParseDatabaseDsn parses the database dsn and extracts user, password, database server, and name.
func (c *Config) ParseDatabaseDsn() {
	if c.options.DatabaseDsn == "" || c.options.DatabaseServer != "" {
		return
	}

	matches := dsnPattern.FindStringSubmatch(c.options.DatabaseDsn)
	names := dsnPattern.SubexpNames()

	for i, match := range matches {
		switch names[i] {
		case "user":
			c.options.DatabaseUser = match
		case "password":
			c.options.DatabasePassword = match
		case "server":
			c.options.DatabaseServer = match
		case "name":
			c.options.DatabaseName = match
		}
	}
}

// DatabaseServer the database server.
func (c *Config) DatabaseServer() string {
	c.ParseDatabaseDsn()

	if c.options.DatabaseServer == "" {
		return "localhost"
	}

	return c.options.DatabaseServer
}

// DatabaseHost the database server host.
func (c *Config) DatabaseHost() string {
	if s := strings.Split(c.DatabaseServer(), ":"); len(s) > 0 {
		return s[0]
	}

	return c.options.DatabaseServer
}

// DatabasePort the database server port.
func (c *Config) DatabasePort() int {
	const defaultPort = 3306

	if s := strings.Split(c.DatabaseServer(), ":"); len(s) != 2 {
		return defaultPort
	} else if port, err := strconv.Atoi(s[1]); err != nil {
		return defaultPort
	} else if port < 1 || port > 65535 {
		return defaultPort
	} else {
		return port
	}
}

// DatabasePortString the database server port as string.
func (c *Config) DatabasePortString() string {
	return strconv.Itoa(c.DatabasePort())
}

// DatabaseName the database schema name.
func (c *Config) DatabaseName() string {
	c.ParseDatabaseDsn()

	if c.options.DatabaseName == "" {
		return "photoprism"
	}

	return c.options.DatabaseName
}

// DatabaseUser returns the database user name.
func (c *Config) DatabaseUser() string {
	c.ParseDatabaseDsn()

	if c.options.DatabaseUser == "" {
		return "photoprism"
	}

	return c.options.DatabaseUser
}

// DatabasePassword returns the database user password.
func (c *Config) DatabasePassword() string {
	c.ParseDatabaseDsn()

	return c.options.DatabasePassword
}

// DatabaseConns returns the maximum number of open connections to the database.
func (c *Config) DatabaseConns() int {
	limit := c.options.DatabaseConns

	if limit <= 0 {
		limit = (runtime.NumCPU() * 2) + 16
	}

	if limit > 1024 {
		limit = 1024
	}

	return limit
}

// DatabaseConnsIdle returns the maximum number of idle connections to the database (equal or less than open).
func (c *Config) DatabaseConnsIdle() int {
	limit := c.options.DatabaseConnsIdle

	if limit <= 0 {
		limit = runtime.NumCPU() + 8
	}

	if limit > c.DatabaseConns() {
		limit = c.DatabaseConns()
	}

	return limit
}

// Db returns the db connection.
func (c *Config) Db() *gorm.DB {
	if c.db == nil {
		log.Fatal("config: database not connected")
	}

	return c.db
}

// CloseDb closes the db connection (if any).
//func (c *Config) CloseDb() error {
//	if c.db != nil {
//		if err := c.db.; err == nil {
//			c.db = nil
//		} else {
//			return err
//		}
//	}
//
//	return nil
//}

// SetDbOptions sets the database collation to unicode if supported.
func (c *Config) SetDbOptions() {
	switch c.DatabaseDriver() {
	case MySQL, MariaDB:
		c.Db().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
	case Postgres:
		// Ignore for now.
	case SQLite:
		// Not required as unicode is default.
	}
}

// InitDb will initialize the database connection and schema.
//func (c *Config) InitDb() {
//	c.SetDbOptions()
//	entity.SetDbProvider(c)
//	entity.MigrateDb(true)
//
//	entity.Admin.InitPassword(c.AdminPassword())
//
//	go entity.SaveErrorMessages()
//}

// InitTestDb drops all tables in the currently configured database and re-creates them.
//func (c *Config) InitTestDb() {
//	c.SetDbOptions()
//	entity.SetDbProvider(c)
//	entity.ResetTestFixtures()
//
//	entity.Admin.InitPassword(c.AdminPassword())
//
//	go entity.SaveErrorMessages()
//}

// connectDb establishes a database connection.
func (c *Config) connectDb() error {
	mutex.Db.Lock()
	defer mutex.Db.Unlock()

	dbDriver := c.DatabaseDriver()
	dbDsn := c.DatabaseDsn()

	if dbDriver == "" {
		return errors.New("config: database driver not specified")
	}

	if dbDsn == "" {
		return errors.New("config: database DSN not specified")
	}

	db, err := gorm.Open(mysql.Open(dbDsn), &gorm.Config{})
	if err != nil || db == nil {
		for i := 1; i <= 12; i++ {
			db, err = gorm.Open(mysql.Open(dbDsn), &gorm.Config{})

			if db != nil && err == nil {
				break
			}

			time.Sleep(5 * time.Second)
		}

		if err != nil || db == nil {
			log.Fatal(err)
		}
	}

	//db.LogMode(false)
	//db.SetLogger(log)
	//
	//db.DB().SetMaxOpenConns(c.DatabaseConns())
	//db.DB().SetMaxIdleConns(c.DatabaseConnsIdle())
	//db.DB().SetConnMaxLifetime(10 * time.Minute)

	c.db = db

	return err
}

// ImportSQL imports a file to the currently configured database.
func (c *Config) ImportSQL(filename string) {
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Error(err)
		return
	}

	statements := strings.Split(string(contents), ";\n")
	q := c.Db().Unscoped()

	for _, stmt := range statements {
		// Skip empty lines and comments
		if len(stmt) < 3 || stmt[0] == '#' || stmt[0] == ';' {
			continue
		}

		var result struct{}

		q.Raw(stmt).Scan(&result)
	}
}
