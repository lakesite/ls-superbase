# ls-superbase #

🏗️ convenience library for datastore configuration in go 🏗️

superbase provides a struct to hold a gorm.DB connection, and datastore config
settings:

```
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
```

## usage ##

```
	DBConfig := make(map[string]*superbase.DBConfig)
	DBConfig["myapp"].Server = "sqlite3"
	DBConfig["myapp"].Path = "myapp.db"
	DBConfig["myapp"].Init()
```

A short example to initialize an sqlite3 db and connection:

### main.go ###

```
package main

import (
	"github.com/lakesite/ls-superbase
)

func main() {
	DBConfig := make(map[string]*superbase.DBConfig)
	DBConfig["myapp"].Server = "sqlite3"
	DBConfig["myapp"].Path = "myapp.db"
	DBConfig["myapp"].Init()
}

```

## dependencies ##

1. 	[gorm](https://github.com/jinzhu/gorm)
2.	[gorm-mysql](https://github.com/jinzhu/gorm/dialects/mysql)
3.  [gorm-postgres](https://github.com/jinzhu/gorm/dialects/postgres)
4.	[gorm-sqlite](https://github.com/jinzhu/gorm/dialects/sqlite)
5.	[gorm-mssql](https://github.com/jinzhu/gorm/dialects/mssql)

## license ##

MIT
