package commands

import (
	"github.com/DKhorkov/golangForPro/phonebook/internal/data"
	"github.com/DKhorkov/golangForPro/phonebook/internal/models"
)

func Search(key string) *models.Entry {
	for i, entry := range data.Entries {
		if entry.Tel == key {
			return &data.Entries[i]
		}
	}

	return nil
}
