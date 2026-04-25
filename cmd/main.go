package main

import (
	"net/http"

	"github.com/dtm/internal/delivery"
)

func main() {

	http.HandleFunc("/landing", delivery.Landing)
	http.HandleFunc("/form", delivery.ServeForm)
	http.HandleFunc("/submit", delivery.Form)
	http.HandleFunc("/reserve", delivery.CustomerReservation)

	http.ListenAndServe(":8080", nil)
}
