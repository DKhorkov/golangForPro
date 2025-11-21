package handlers

import (
	"log"
	"net/http"
)

func GetFileHandler(filepath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

		http.ServeFile(w, r, filepath)
	}
}
