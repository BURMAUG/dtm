package delivery

import (
	"log"
	"net/http"
)

func Landing(w http.ResponseWriter, r *http.Request) {
}

// serve the form /form
func ServeForm(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
}

// submit the form with data /reserve
func Form(w http.ResponseWriter, r *http.Request) {
}

// get and displays customer information based on the id being sent.
func CustomerReservation(w http.ResponseWriter, r *http.Request) {

}
