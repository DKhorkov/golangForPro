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
	reverseKey = "reverse"
)

func ListHandler(pb interfaces.PhoneBook) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

			return
		}

		reverseStr := strings.TrimSpace(r.URL.Query().Get(reverseKey))
		reverse, _ := strconv.ParseBool(reverseStr)

		entries, err := pb.List(reverse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err = json.NewEncoder(w).Encode(entries); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}
