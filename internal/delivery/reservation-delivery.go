package delivery

import (
	"html/template"
	"log"
	"net/http"

	"github.com/dtm/internal/usecase"
)

const (
	indexHTML       = "../template/index.gohtml"
	reservationHTML = "../template/reservation.gohtml"
	formHTML        = "../template/form.gohtml"
	thanksHTML      = "../template/thanks.gohtml"
)

var customerReservation = usecase.CustomerReservationUsecase{}

func Landing(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	tmpl, err := template.ParseFiles(indexHTML)
	if r.Method != http.MethodGet || err != nil {
		log.Print(err)
		http.Error(w, "Sorry", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// serve the form /form
func ServeForm(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	tmpl, err := template.ParseFiles(formHTML)
	if r.Method != http.MethodGet || err != nil {
		log.Print(r.Method)
		log.Print("form error ", err)
		http.Error(w, "Sorry", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
	// talk to usecase layer
}

// submit the form with data /reserve
func Form(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.Method != http.MethodPost {
		log.Print(r.Method)
		http.Error(w, "Sorry", http.StatusInternalServerError)
		return
	}

	// talk to usecase layer
	customerReservation.MakeReservation(r.Context(), r)
	tmpl, err := template.ParseFiles(thanksHTML)
	if err != nil {
		log.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}

// get and displays customer information based on the id being sent.
func CustomerReservation(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	tmpl, err := template.ParseFiles(reservationHTML)
	if r.Method != http.MethodGet || err != nil {
		http.Error(w, "Sorry", http.StatusInternalServerError)
		return
	}

	customerReservation.GetCustomerReservation(r.Context(), r)
	tmpl.Execute(w, nil)
	// talk to usecase layer
}
