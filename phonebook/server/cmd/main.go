package main

import (
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/app"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/config"
	controllers "github.com/DKhorkov/golangForPro/phonebook/server/internal/controllers/http"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/filepath"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/repositories"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/usecases"
	"github.com/DKhorkov/libs/loadenv"
)

func main() {
	// Инициализируем переменные окружения для дальнейшего считывания в конфиге:
	loadenv.Init()

	cfg := config.New()

	fp, err := filepath.Get(cfg.Filepath.Path, cfg.Filepath.Source)
	if err != nil {
		panic(err)
	}

	entriesRepository, err := repositories.New(fp)
	if err != nil {
		panic(err)
	}

	u, err := usecases.New(entriesRepository)
	if err != nil {
		panic(err)
	}

	c := controllers.New(cfg.HTTP, cfg.CORS, u)

	application := app.New(c)
	application.Run()
}
