package models

import "github.com/DKhorkov/golangForPro/phonebook/internal/commands"

type Command struct {
	Name   commands.CommandName `json:"name"`
	Params []string             `json:"params"`
}
