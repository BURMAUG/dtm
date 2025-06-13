package domain

import (
	"time"

	"github.com/google/uuid"
)

type Customer any // revisit wha this is and why it is here

type Addr struct {
	AddressId uuid.UUID `gorm:"primarykey;type:uuid;column:address_id"`
	Address   string    `gorm:"not null;column:address"`
	City      string    `gorm:"not null;column:city"`
	State     string    `gorm:"not null;column:state"`
	Zip       string    `gorm:"not null;column:zip"`
}

type CustomerInfo struct {
	CustomerId     uuid.UUID `gorm:"primarykey;type:uuid;column:customer_id"`
	Name           string    `gorm:"not null;column:name"`
	Email          string    `gorm:"not null;column:email"`
	Phone          string    `gorm:"not null;column:phone"`
	House          bool      `gorm:"not null;column:is_house"`
	NumberOfRooms  int       `gorm:"not null;column:number_of_rooms"`
	PickAddress    Addr      `gorm:"embedded;embeddedPrefix:pick_up_"`
	DropOffAddress Addr      `gorm:"embedded;embeddedPrefix:drop_off-"`
	Date           time.Time `gorm:"not null;column:date_and_time"`
}
