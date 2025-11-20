package interfaces

import "github.com/DKhorkov/golangForPro/phonebook/server/internal/models"

type ReadWriter interface {
	Read() ([]models.Entry, error)
	Write(entries []models.Entry) error
}
