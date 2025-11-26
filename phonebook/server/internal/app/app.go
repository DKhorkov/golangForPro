package app

import (
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	controller interfaces.Controller
}

func (application *App) Run() {
	// Launch asynchronous for graceful shutdown purpose:
	go application.controller.Run()

	// Graceful shutdown. When system signal will be received, signal.Notify function will write it to channel.
	// After this event, main goroutine will be unblocked (<-stopChannel blocks it) and application will be
	// gracefully stopped:
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)
	<-stopChannel
	application.controller.Stop()
}

func New(controller interfaces.Controller) *App {
	return &App{
		controller: controller,
	}
}
