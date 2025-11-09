package commands

import "github.com/DKhorkov/golangForPro/phonebook/internal/models"

func List() []models.Entry {
	return entries
}
