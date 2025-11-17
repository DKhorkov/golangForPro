package interfaces

import "github.com/DKhorkov/golangForPro/phonebook/internal/models"

type PhoneBook interface {
	List(reverse bool) ([]models.Entry, error)
	Insert(entry models.Entry) error
	Delete(key string) error
	Search(key string) (*models.Entry, error)
}
