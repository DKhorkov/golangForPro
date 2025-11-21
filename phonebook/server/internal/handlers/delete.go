package handlers

import (
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/validation"
	"log"
	"net/http"
	"strings"
)

const (
	deleteKey = "key"
)

func DeleteHandler(pb interfaces.PhoneBook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

		key := r.URL.Query().Get(searchKey)
		if !strings.HasPrefix(key, "+") {
			key = "+" + key
		}

		if !validation.ValidatePhone(key) {
			http.Error(w, "Not a valid Phone. Phone should be like \"+7 (911) 258-01-62\"\n", http.StatusBadRequest)

			return
		}

		if err := pb.Delete(key); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}
