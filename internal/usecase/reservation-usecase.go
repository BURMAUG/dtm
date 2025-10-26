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
	GetCustomerReservation(ctx context.Context, r *http.Request) // why is this get? it does not return anything
	MakeReservation(w http.ResponseWriter, r *http.Request)      // can return void
}

type CustomerReservationUsecase struct{}

func (c *CustomerReservationUsecase) MakeReservation(ctx context.Context, r *http.Request) {
	customer, err := extractCustomerData(r)
	if err != nil {
		log.Print(err)
		return
	}

	email = EmailUsecase{CustomerInfo: *customer}

	customerRepsitory = &repository.CustomerInfo{}
	err = customerRepsitory.SaveCustomerInfo(ctx, customer)
	if err != nil {
		log.Print(err)
		return
	}

	// non blocking
	go func() {
		email.SendEmail()
	}()

	// non blocking
	go func() {
		email.SendEmail()
	}()
}

func (c *CustomerReservationUsecase) GetCustomerReservation(ctx context.Context, r *http.Request) {
	// extract the query id here
	id, err := uuid.Parse(r.URL.Query().Get("id"))
	check(err)

	// return customer data
	customerRepsitory.FindCustomerById(ctx, id)
}

func extractCustomerData(r *http.Request) (*domain.CustomerInfo, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	pickUp, drop, err := extractAddress(r)
	go func() { check(err) }()

	id, err := uuid.NewUUID()
	go func() { check(err) }()

	// time, err := time.Parse(time.RFC3339, r.FormValue("date"))
	// check(err)
	time := time.Now() //Todo() this has to change

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
	wg.Done()
	wg.Wait()
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
