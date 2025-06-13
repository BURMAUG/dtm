package config

import (
	"errors"
	"log"

	"github.com/dtm/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	dsn := "host=localhost user=myuser password=password dbname=mydb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(errors.New("failed to open db connection: "), err)
	}
	// Drop all tables
	err = db.Migrator().DropTable(
		&domain.CustomerInfo{}, // add all your domain models here
	)
	if err != nil {
		log.Fatal("failed to drop tables: ", err)
	}
	err = db.Migrator().DropTable(
		&domain.Addr{}, // add all your domain models here
	)
	if err != nil {
		log.Fatal("failed to drop tables: ", err)
	}

	// Create all tables
	err = db.AutoMigrate(
		&domain.CustomerInfo{}, // add them again here
	)
	if err != nil {
		log.Fatal("failed to migrate tables: ", err)
	}

}

func Conn() *gorm.DB {
	return db
}
