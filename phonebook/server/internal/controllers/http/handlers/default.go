package handlers

import (
	"log"
	"net/http"
)

// swagger:route POST / DefaultHandler OK
// Default Handler for everything that is not a match.
// Works with all HTTP methods
//
// responses:
//  200: OK
//  500: InternalServerError

// DefaultHandler is for handling everything that is not a match
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("Thanks for visiting!\n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
