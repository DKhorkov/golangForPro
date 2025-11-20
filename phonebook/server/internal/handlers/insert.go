package handlers

import (
	"encoding/json"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/models"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/validation"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	nameKey    = "name"
	surnameKey = "surname"
	phoneKey   = "phone"
)

func InsertHandler(pb interfaces.PhoneBook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

		name := strings.TrimSpace(r.URL.Query().Get(nameKey))
		if !validation.ValidateNameSurname(name) {
			http.Error(w, "Not a valid Name. Name should be like \"Илья\" or \"Ilya\"\n", http.StatusBadRequest)

			return
		}

		surname := strings.TrimSpace(r.URL.Query().Get(surnameKey))
		if !validation.ValidateNameSurname(surname) {
			http.Error(w, "Not a valid Surname. Surname should be like \"Романов\" or \"Romanov\"\n", http.StatusBadRequest)

			return
		}

		phone := "+" + strings.TrimSpace(r.URL.Query().Get(phoneKey))
		if !validation.ValidatePhone(phone) {
			http.Error(w, "Not a valid Phone. Phone should be like \"+7 (911) 258-01-62\"\n", http.StatusBadRequest)

			return
		}

		entry := models.Entry{
			Name:       name,
			Surname:    surname,
			Phone:      phone,
			LastAccess: time.Now(),
		}

		if err := pb.Insert(entry); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := json.NewEncoder(w).Encode(entry); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}
