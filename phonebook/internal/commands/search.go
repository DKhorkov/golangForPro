package commands

import (
	"github.com/DKhorkov/golangForPro/phonebook/internal/models"
)

func Search(key string) *models.Entry {
	for i, entry := range entries {
		if entry.Surname == key {
			return &entries[i]
		}
	}

	return nil
}
