package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, _ = gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})

func DBConnection() *gorm.DB {
	return db
}
