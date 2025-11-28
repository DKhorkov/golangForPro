package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	customErrors "github.com/DKhorkov/golangForPro/phonebook/server/internal/errors"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/validation"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

const (
	SearchKey = "key"
)

// swagger:route GET /entries/{key} SearchEntry searchKey
// Returns searched entry
//
// responses:
//	200: Entry
//  400: BadRequest
//  404: NotFound
//  500: InternalServerError

// SearchHandler returns searched entry
func SearchHandler(u interfaces.UseCases) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

		key := mux.Vars(r)[SearchKey]
		if !strings.HasPrefix(key, "+") {
			key = "+" + key
		}

		if !validation.ValidatePhone(key) {
			http.Error(
				w,
				fmt.Sprintf(
					"Not a valid Phone \"%s\". Phone should be like \"+7 (911) 258-01-62\"\n",
					key,
				),
				http.StatusBadRequest,
			)

			return
		}

		entry, err := u.Search(key)
		switch {
		case errors.Is(err, customErrors.ErrNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)

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
