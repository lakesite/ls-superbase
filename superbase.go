package superbase

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// DBConfig holds a database configuration and gorm.DB connection.
type DBConfig struct {
	Connection *gorm.DB
	Server     string
	Port       string
	Database   string
	User       string
	Password   string
	Driver     string
	Path       string
}

// Init initalizes our database connection based on the driver.
func (db *DBConfig) Init() {
	if db.Driver == "sqlite3" {
		db.Connection, _ = gorm.Open("sqlite3", db.Path)
	} else {
		// handle connections with other drivers
	}
}
