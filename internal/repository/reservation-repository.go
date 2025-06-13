package repository

import (
	"context"

	"github.com/dtm/internal/config"
	"github.com/dtm/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB = config.Conn()

type Customer interface {
	FindCustomerById(ctx context.Context, id uuid.UUID)
	SaveCustomerInfo(ctx context.Context, customerInfo *domain.CustomerInfo) error
}

type CustomerData struct{}

func (c *CustomerData) FindCustomerById(ctx context.Context, id uuid.UUID) {
	db.WithContext(ctx).First(&domain.CustomerInfo{}, "customer_id = ? ", id)
}

func (c *CustomerData) SaveCustomerInfo(ctx context.Context, customerInfo *domain.CustomerInfo) error {
	if err := db.WithContext(ctx).Save(customerInfo).Error; err != nil {
		db.AddError(err)
		return err
	}
	return nil

}
