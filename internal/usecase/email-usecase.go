package usecase

import (
	"log"

	"github.com/dtm/internal/domain"
	"gopkg.in/gomail.v2"
)

type Email interface {
	SendEmail()
	formatAddress() (string, error)
}

type EmailUsecase struct {
	CustomerInfo domain.CustomerInfo
}

func (e *EmailUsecase) SendEmail() {
	m := gomail.NewMessage()
	m.SetHeader("From", e.CustomerInfo.Email)
	m.SetHeader("To", "jeftadjg@gmail.com")
	m.SetHeader("Subject", "DTM: Moving Reservation")

	p, da, _ := e.formatAddress()
	message := "Moving alert from Customer: " + e.CustomerInfo.Name + ".\n\nPickup address: " + p + "\n\nDrop off Address: " + da
	m.SetBody("text/plain", message)

	d := gomail.NewDialer("smtp.gmail.com", 587, e.CustomerInfo.Email, "oenourfgsfdqxtwa")
	if err := d.DialAndSend(m); err != nil {
		log.Print("There was an  error sending Email: ", err)
		return
	}
	log.Print("@@@@@@@@@@@@Email Sent Customer@@@@@@@@@@@@@@@")
}

// format
func (e *EmailUsecase) formatAddress() (string, string, error) {
	p := e.CustomerInfo.PickAddress
	d := e.CustomerInfo.DropOffAddress
	pickupadd := p.Address + " " + p.City + " " + p.State + " " + p.Zip
	droppoff := d.Address + " " + d.City + " " + d.State + " " + d.Zip
	return pickupadd, droppoff, nil
}
