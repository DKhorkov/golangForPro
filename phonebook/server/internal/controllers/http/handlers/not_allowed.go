package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func NotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, "Method:", r.Method)

	w.WriteHeader(http.StatusMethodNotAllowed)

	_, err := w.Write([]byte(fmt.Sprintf("Method \"%s\" not allowed for URL \"%s\"!\n", r.Method, r.URL.Path)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
