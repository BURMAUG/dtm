package main

import (
	"net/http"

	_ "github.com/dtm/internal/config"
	"github.com/dtm/internal/delivery"
)

func main() {
	http.HandleFunc("/", delivery.Landing)       // home index.gohtm
	http.HandleFunc("/form", delivery.ServeForm) //form.gohtml
	http.HandleFunc("/submit", delivery.Form)
	http.HandleFunc("/reserve", delivery.CustomerReservation)

	//NotFound test
	http.HandleFunc("", http.NotFound) //this should be deleted as it is not needed at all!

	http.ListenAndServe(":8080", nil)
}
