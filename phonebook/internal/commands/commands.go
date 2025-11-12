package commands

type CommandName string

const (
	CommandSearch CommandName = "search"
	CommandList   CommandName = "list"
	CommandInsert CommandName = "insert"
	CommandDelete CommandName = "delete"
)
