package config

import (
	"errors"
	"log"

	"github.com/dtm/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "host=localhost user=myuser password=password dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(errors.New("failed to open db connection: "), err)
	}
	db.AutoMigrate(&domain.CustomerInfo{})
}

func Conn() *gorm.DB {
	return db
}
