package main

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/config"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/filepath"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/handlers"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/phonebook"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/readwriters"
	"github.com/DKhorkov/libs/loadenv"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	// Инициализируем переменные окружения для дальнейшего считывания в конфиге:
	loadenv.Init()

	cfg := config.New()

	fp, err := filepath.Get(cfg.Filepath.Path, cfg.Filepath.Source)
	if err != nil {
		panic(err)
	}

	rw, err := readwriters.New(fp)
	if err != nil {
		panic(err)
	}

	pb, err := phonebook.New(rw)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handlers.DefaultHandler))
	mux.Handle("/list", handlers.ListHandler(pb))
	mux.Handle("/insert", handlers.InsertHandler(pb))
	mux.Handle("/search", handlers.SearchHandler(pb))
	mux.Handle("/delete", handlers.DeleteHandler(pb))
	mux.Handle("/status", handlers.StatusHandler(pb))
	mux.Handle("/getFile", handlers.GetFileHandler(fp))

	httpHandler := cors.New(
		cors.Options{
			AllowedOrigins:   cfg.CORS.AllowedOrigins,
			AllowedMethods:   cfg.CORS.AllowedMethods,
			AllowedHeaders:   cfg.CORS.AllowedHeaders,
			MaxAge:           cfg.CORS.MaxAge,
			AllowCredentials: cfg.CORS.AllowCredentials,
		},
	).Handler(mux)

	addr := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      httpHandler,
		IdleTimeout:  cfg.HTTP.IdleTimeout,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
	}

	fmt.Println("Ready to serve at ", addr)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
