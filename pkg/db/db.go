package db

import (
	"fmt"

	"github.com/go-baselib/go-plugin/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(db *config.DB) (*gorm.DB, error) {
	if db == nil {
		return nil, fmt.Errorf("db conf invalid")
	}

	var dial gorm.Dialector
	switch db.Type {
	case "sqlite", "":
		dial = sqlite.Open(db.DSN)
	case "mysql":
		dial = mysql.Open(db.DSN)
	case "postgres":
		dial = postgres.Open(db.DSN)
	}

	return gorm.Open(dial, &gorm.Config{})

}
