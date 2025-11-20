package models

import (
	"fmt"
	"time"
)

type Entry struct {
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Phone      string    `json:"phone"`
	LastAccess time.Time `json:"lastAccess"`
}

func (e *Entry) View() string {
	return fmt.Sprintf(
		"Name: %s\nSurname: %s\nPhone: %s\nLastAccess: %s\n",
		e.Name,
		e.Surname,
		e.Phone,
		e.LastAccess.Format(time.RFC1123),
	)
}
