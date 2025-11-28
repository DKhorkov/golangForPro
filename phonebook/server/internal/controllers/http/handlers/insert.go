package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	customErrors "github.com/DKhorkov/golangForPro/phonebook/server/internal/errors"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/models"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/validation"
	"io"
	"log"
	"net/http"
	"strings"
)

// swagger:route POST /entries InsertEntry insertEntryInput
// Create a new Entry
//
// responses:
//	200: Entry
//  400: BadRequest
//  500: InternalServerError

// InsertHandler is for adding a new entry
func InsertHandler(u interfaces.UseCases) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		var entry models.Entry
		if err = json.Unmarshal(data, &entry); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		name := strings.TrimSpace(entry.Name)
		if !validation.ValidateNameSurname(name) {
			http.Error(
				w,
				fmt.Sprintf(
					"Not a valid Name \"%s\". Surname should be like \"Илья\" or \"Ilya\"\n",
					name,
				),
				http.StatusBadRequest,
			)

			return
		}

		surname := strings.TrimSpace(entry.Surname)
		if !validation.ValidateNameSurname(surname) {
			http.Error(
				w,
				fmt.Sprintf(
					"Not a valid Surname \"%s\". Surname should be like \"Романов\" or \"Romanov\"\n",
					surname,
				),
				http.StatusBadRequest,
			)

			return
		}

		phone := strings.TrimSpace(entry.Phone)
		if !strings.HasPrefix(phone, "+") {
			phone = "+" + phone
		}

		if !validation.ValidatePhone(phone) {
			http.Error(
				w,
				fmt.Sprintf(
					"Not a valid Phone \"%s\". Phone should be like \"+7 (911) 258-01-62\"\n",
					phone,
				),
				http.StatusBadRequest,
			)

			return
		}

		err = u.Insert(entry)
		switch {
		case errors.Is(err, customErrors.ErrAlreadyExists):
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		case err != nil:
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		if err = json.NewEncoder(w).Encode(entry); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
