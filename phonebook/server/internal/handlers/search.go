package handlers

import (
	"encoding/json"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/validation"
	"log"
	"net/http"
	"strings"
)

const (
	searchKey = "key"
)

func SearchHandler(pb interfaces.PhoneBook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

		key := "+" + strings.TrimSpace(r.URL.Query().Get(searchKey))
		if !validation.ValidatePhone(key) {
			http.Error(w, "Not a valid Phone. Phone should be like \"+7 (911) 258-01-62\"\n", http.StatusBadRequest)

			return
		}

		entry, err := pb.Search(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err = json.NewEncoder(w).Encode(entry); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}
