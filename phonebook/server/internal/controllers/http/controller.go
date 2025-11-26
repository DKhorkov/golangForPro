package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/config"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/controllers/http/handlers"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/middlewares"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"net/http"
)

type Controller struct {
	server *http.Server
	host   string
	port   int
}

func New(
	httpConfig config.HTTPConfig,
	corsConfig config.CORSConfig,
	useCases interfaces.UseCases,
) *Controller {
	rootMux := mux.NewRouter()
	rootMux.NotFoundHandler = http.HandlerFunc(handlers.DefaultHandler)
	rootMux.MethodNotAllowedHandler = http.HandlerFunc(handlers.NotAllowedHandler)
	rootMux.Use(middlewares.MetricsMiddleware)

	getMux := rootMux.Methods(http.MethodGet).Subrouter()
	getMux.Handle("/metrics", promhttp.Handler())
	getMux.Handle("/entries", handlers.ListHandler(useCases))
	getMux.Handle(fmt.Sprintf("/entries/{%s}", handlers.SearchKey), handlers.SearchHandler(useCases))

	postMux := rootMux.Methods(http.MethodPost).Subrouter()
	postMux.Handle("/entries", handlers.InsertHandler(useCases))

	deleteMux := rootMux.Methods(http.MethodDelete).Subrouter()
	deleteMux.Handle(fmt.Sprintf("/entries/{%s}", handlers.DeleteKey), handlers.DeleteHandler(useCases))

	httpHandler := cors.New(
		cors.Options{
			AllowedOrigins:   corsConfig.AllowedOrigins,
			AllowedMethods:   corsConfig.AllowedMethods,
			AllowedHeaders:   corsConfig.AllowedHeaders,
			MaxAge:           corsConfig.MaxAge,
			AllowCredentials: corsConfig.AllowCredentials,
		},
	).Handler(rootMux)

	addr := fmt.Sprintf("%s:%d", httpConfig.Host, httpConfig.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      httpHandler,
		IdleTimeout:  httpConfig.IdleTimeout,
		ReadTimeout:  httpConfig.ReadTimeout,
		WriteTimeout: httpConfig.WriteTimeout,
	}

	return &Controller{
		server: server,
		host:   httpConfig.Host,
		port:   httpConfig.Port,
	}
}

func (c *Controller) Run() {
	addr := fmt.Sprintf("%s:%d", c.host, c.port)
	fmt.Println("Ready to serve at ", addr)

	if err := c.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("HTTP server error: %v\n", err)
	}

	fmt.Println("Stopped serving new connections.")
}

func (c *Controller) Stop() {
	// Stops accepting new requests and processes already received requests:
	err := c.server.Shutdown(context.Background())
	if err != nil {
		fmt.Printf("HTTP shutdown error: %v\n", err)
	}

	fmt.Println("Graceful shutdown completed.")
}
