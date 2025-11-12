package interfaces

import "github.com/DKhorkov/golangForPro/phonebook/internal/models"

type PhoneBook interface {
	Execute(command models.Command) error
}
