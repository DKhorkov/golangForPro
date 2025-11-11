package data

import "github.com/DKhorkov/golangForPro/phonebook/internal/models"

const (
	nameLength    = 7
	surnameLength = 10

	entriesCount = 100
)

var Entries []models.Entry

func init() {
	for range entriesCount {
		Entries = append(
			Entries, models.Entry{
				Name:    GenerateRandomString(nameLength),
				Surname: GenerateRandomString(surnameLength),
				Tel:     GeneratePhoneNumber(),
			},
		)
	}
}
