package models

type CommandName string

const (
	CommandSearch CommandName = "search"
	CommandList   CommandName = "list"
)

type Command struct {
	Name   CommandName `json:"name"`
	Params []string    `json:"params"`
}
