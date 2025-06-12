package repository

import (
	"context"

	"github.com/dtb/internal/domain"
	"github.com/google/uuid"
)

type CustomerRepository interface {
	FindCustomerById(id uuid.UUID)
	SaveCustomerInfo(ctx context.Context, customerInfo *domain.CustomerInfo) error
}
