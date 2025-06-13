package repository

import (
	"context"
	"log"

	"github.com/dtm/internal/config"
	"github.com/dtm/internal/domain"
	"github.com/google/uuid"
)

var db = config.Conn()

type Customer interface {
	FindCustomerById(ctx context.Context, id uuid.UUID)
	SaveCustomerInfo(ctx context.Context, customerInfo *domain.CustomerInfo) error
}

type CustomerInfo struct{}

func (c *CustomerInfo) FindCustomerById(ctx context.Context, id uuid.UUID) {
	db.WithContext(ctx).First(&domain.CustomerInfo{}, "customer_id = ? ", id)
}

func (c *CustomerInfo) SaveCustomerInfo(ctx context.Context, customerInfo *domain.CustomerInfo) error {
	log.Print("@@@@@@@@@@@@@@@@@@@Customer", customerInfo)
	log.Print("@@@@@@@@@@DB", db)

	if err := db.WithContext(ctx).Create(customerInfo).Error; err != nil {
		log.Print("DB Errorr @@@@@@@ ", err)
		return err
	}
	return nil

}
