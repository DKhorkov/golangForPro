package commands

import (
	"github.com/DKhorkov/golangForPro/phonebook/internal/data"
	"github.com/DKhorkov/golangForPro/phonebook/internal/models"
)

func List() []models.Entry {
	return data.Entries
}
