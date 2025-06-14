package main

import (
	"net/http"

	_ "github.com/dtm/internal/config"
	"github.com/dtm/internal/delivery"
)

func main() {
	http.HandleFunc("/landing", delivery.Landing) // home index.gohtm
	http.HandleFunc("/form", delivery.ServeForm)  //form.gohtml
	http.HandleFunc("/submit", delivery.Form)
	http.HandleFunc("/reserve", delivery.CustomerReservation)

	http.ListenAndServe(":8080", nil)
}
