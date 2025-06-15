package usecase

import (
	"context"
	"log"
	"net/http"
	"sync"
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
	email = EmailUsecase{CustomerInfo: *customer}
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
	// non blocking
	go func() { email.SendAdminEmail() }()
	// non blocking
	go func() { email.SendCustomerEmail() }()
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
	go func() { check(err) }()

	id, err := uuid.NewUUID()
	go func() { check(err) }()

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

	var pickUpAddr *domain.Addr
	var dropOffAddress *domain.Addr

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		pickUpAddr = getAddress("p", r)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		dropOffAddress = getAddress("p", r)
		wg.Done()
	}()
	wg.Wait()

	return *pickUpAddr, *dropOffAddress, nil
}

func getAddress(prefix string, r *http.Request) *domain.Addr {
	id, err := uuid.NewUUID()
	go func() { check(err) }()
	line := r.FormValue(prefix + "addr")
	if isEmpty(line) {
		log.Print(line)
	}

	city := r.FormValue(prefix + "city")
	if isEmpty(city) {
		log.Print(city)
	}

	state := r.FormValue(prefix + "state")
	if isEmpty(state) {
		log.Print(state)
	}

	zip := r.FormValue(prefix + "zip")
	if isEmpty(zip) {
		log.Print(zip)
	}
	return &domain.Addr{
		AddressId: id,
		Address:   line,
		City:      city,
		State:     state,
		Zip:       zip,
	}
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
