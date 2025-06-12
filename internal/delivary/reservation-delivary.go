package delivary

import (
	"log"
	"net/http"
)

func Landing(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
}
