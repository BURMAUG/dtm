package usecase

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dtm/internal/domain"
	"github.com/dtm/internal/repository"
	"github.com/google/uuid"
)

var customerRepsitory repository.Customer
var email EmailUsecase

type Resevation interface {
	GetCustomerReservation(ctx context.Context, r *http.Request)
	MakeReservation(w http.ResponseWriter, r *http.Request)
}

type CustomerReservationUsecase struct {
}

func (c *CustomerReservationUsecase) MakeReservation(ctx context.Context, r *http.Request) {
	customer, err := extractCustomerData(r)
	if err != nil {
		log.Print(err)
		return
	}
	customerRepsitory = &repository.CustomerInfo{}
	err = customerRepsitory.SaveCustomerInfo(ctx, customer)
	if err != nil {
		log.Print(err)
		return
	}
	//send us email first
	email = EmailUsecase{CustomerInfo: *customer}
	email.SendAdminEmail()

}

func (c *CustomerReservationUsecase) GetCustomerReservation(ctx context.Context, r *http.Request) {
	// extract the query id here
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	check(err)

	// return customer data
	customerRepsitory.FindCustomerById(ctx, id)
}

func extractCustomerData(r *http.Request) (*domain.CustomerInfo, error) {
	pickUp, drop, err := extractAddress(r)
	check(err)

	id, err := uuid.NewUUID()
	check(err)

	// time, err := time.Parse(time.RFC3339, r.FormValue("date"))
	// check(err)
	time := time.Now()
	customer := &domain.CustomerInfo{
		CustomerId:     id,
		Name:           r.FormValue("name"),
		Email:          r.FormValue("email"),
		Phone:          r.FormValue("phone"),
		House:          true,
		NumberOfRooms:  2,
		PickAddress:    pickUp,
		DropOffAddress: drop,
		Date:           time,
	}
	return customer, nil
}

func extractAddress(r *http.Request) (domain.Addr, domain.Addr, error) {
	id, err := uuid.NewUUID()
	check(err)

	line := r.FormValue("paddr")
	if isEmpty(line) {
		log.Print(line)
	}

	city := r.FormValue("pcity")
	if isEmpty(city) {
		log.Print(city)
	}

	state := r.FormValue("pstate")
	if isEmpty(state) {
		log.Print(state)
	}

	zip := r.FormValue("pzip")
	if isEmpty(zip) {
		log.Print(zip)
	}

	pickUpAddr := domain.Addr{
		AddressId: id,
		Address:   line,
		City:      city,
		State:     state,
		Zip:       zip,
	}

	line = r.FormValue("daddr")
	if isEmpty(line) {
		log.Print(line)
	}

	city = r.FormValue("dcity")
	if isEmpty(city) {
		log.Print(city)
	}

	state = r.FormValue("dstate")
	if isEmpty(state) {
		log.Print(state)
	}

	zip = r.FormValue("dzip")
	if isEmpty(zip) {
		log.Print(zip)
	}

	dropOffAddress := domain.Addr{
		AddressId: id,
		Address:   line,
		City:      city,
		State:     state,
		Zip:       zip,
	}

	return pickUpAddr, dropOffAddress, nil
}
func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func isEmpty(str string) bool {
	if len(str) < 1 || str == "" {
		return true
	}
	return false
}
