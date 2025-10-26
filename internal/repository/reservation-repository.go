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
	FindCustomerById(ctx context.Context, id uuid.UUID) (CustomerInfo, error)
	SaveCustomerInfo(ctx context.Context, customerInfo *domain.CustomerInfo) error
}

type CustomerInfo struct{}

// FindCustomerById finds a customer by an id I think if you find something
// it should be return
func (c *CustomerInfo) FindCustomerById(ctx context.Context, id uuid.UUID) (CustomerInfo, error) {
	db.WithContext(ctx).First(&domain.CustomerInfo{}, "customer_id = ? ", id)
	// todo return CustomerInfo

	//todo
	return CustomerInfo{}, nil
}

// SaveCustomerInfo public function that takes customerInfo and saves it in the database of choice
// returns an error if any
func (c *CustomerInfo) SaveCustomerInfo(ctx context.Context, customerInfo *domain.CustomerInfo) error {
	log.Print("@@@@@@@@@@@@@@@@@@@Customer", customerInfo)
	log.Print("@@@@@@@@@@DB", db)

	if err := db.WithContext(ctx).Create(customerInfo).Error; err != nil {
		log.Print("DB Errorr @@@@@@@ ", err)
		return err
	}
	return nil
}
