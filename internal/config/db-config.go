package config

import (
	"context"
	"errors"
	"log"

	"github.com/dtb/internal/domain"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Database interface {
	GetCustomerById(id uuid.UUID) (*domain.CustomerInfo, error)
	SaveCustomerToDb(ctx context.Context, customer *domain.CustomerInfo) error
}

type DbConfig struct{}

func init() {
	dsn := "host=localhost user=myuser password=password dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(errors.New("failed to open db connection: "), err)
	}
	db.AutoMigrate(&domain.CustomerInfo{})
}
