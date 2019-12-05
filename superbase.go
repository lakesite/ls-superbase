package superbase

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// DBConfig holds a database configuration and gorm.DB connection.
// Meta is used for database related key values that apps might use.
type DBConfig struct {
	Connection *gorm.DB
	Server     string
	Port       string
	Database   string
	User       string
	Password   string
	Driver     string
	Path       string
	Meta       map[string]string
}

// Init initalizes our database connection based on the driver.
func (db *DBConfig) Init() {
	if db.Driver == "sqlite3" {
		db.Connection, _ = gorm.Open("sqlite3", db.Path)
	} else {
		// handle connections with other drivers
	}
	db.Meta = make(map[string]string)
}
