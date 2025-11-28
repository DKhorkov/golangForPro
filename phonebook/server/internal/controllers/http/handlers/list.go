package handlers

import (
	"encoding/json"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	ReverseKey = "reverse"
)

// swagger:route GET /entries ListEntries reverseKey
// Returns all entries
//
// responses:
//	200: Entries
//  500: InternalServerError

// ListHandler returns all entries
func ListHandler(u interfaces.UseCases) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

		reverseStr := strings.TrimSpace(r.URL.Query().Get(ReverseKey))
		reverse, _ := strconv.ParseBool(reverseStr)

		entries, err := u.List(reverse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		if err = json.NewEncoder(w).Encode(entries); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
