package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Default *gorm.DB

// Init all database's connections.
// Called once at startup.
func Init() (err error) {
	dsn := os.Getenv("DB_DNS")

	Default, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return
}
