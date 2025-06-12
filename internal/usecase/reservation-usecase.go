package usecase

import (
	"net/http"

	"github.com/dtb/internal/repository"
)

type Resevation interface {
	FetchCustomerReservationById(r *http.Request)
	SaveCustomerInfo(w http.ResponseWriter, r *http.Request)
}

type Reservations struct {
	customerRepository repository.CustomerRepository
}
