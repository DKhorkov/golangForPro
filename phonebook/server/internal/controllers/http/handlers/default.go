package handlers

import (
	"log"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("Thanks for visiting!\n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
