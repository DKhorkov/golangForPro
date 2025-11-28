package handlers

import (
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
	DeleteKey = "key"
)

// swagger:route DELETE /entries/{key} DeleteEntry deleteKey
// Delete an entry given it key.
//
// responses:
//  200: OK
//  400: BadRequest
//  404: NotFound
//  500: InternalServerError

// DeleteHandler is for deleting users based on provided key
func DeleteHandler(u interfaces.UseCases) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

		key := mux.Vars(r)[DeleteKey]
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

		err := u.Delete(key)
		switch {
		case errors.Is(err, customErrors.ErrNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)

			return
		case err != nil:
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
