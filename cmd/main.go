package main

import (
	"net/http"

	"github.com/dtb/internal/delivary"
)

func main() {
	http.HandleFunc("/landing", delivary.Landing)

	http.ListenAndServe(":8080", nil)
}
