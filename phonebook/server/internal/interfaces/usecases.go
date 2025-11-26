package interfaces

import "github.com/DKhorkov/golangForPro/phonebook/server/internal/models"

type UseCases interface {
	List(reverse bool) ([]models.Entry, error)
	Insert(entry models.Entry) error
	Delete(key string) error
	Search(key string) (*models.Entry, error)
}
