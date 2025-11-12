package interfaces

import "github.com/DKhorkov/golangForPro/phonebook/internal/models"

type ReadWriter interface {
	Read() ([]models.Entry, error)
	Write(entries []models.Entry) error
}
