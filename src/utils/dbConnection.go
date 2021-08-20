package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
}
