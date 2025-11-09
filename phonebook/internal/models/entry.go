package models

import "fmt"

type Entry struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Tel     string `json:"tel"`
}

func (e *Entry) View() string {
	return fmt.Sprintf(
		"Name: %s\nSurname: %s\nPhone:%s\n",
		e.Name,
		e.Surname,
		e.Tel,
	)
}
