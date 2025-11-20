package handlers

import (
	"encoding/json"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"log"
	"net/http"
)

func StatusHandler(pb interfaces.PhoneBook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

		entries, err := pb.List(false)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err = json.NewEncoder(w).Encode(len(entries)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}
