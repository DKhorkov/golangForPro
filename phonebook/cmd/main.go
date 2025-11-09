package main

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/internal/args"
	"github.com/DKhorkov/golangForPro/phonebook/internal/commands"
	"github.com/DKhorkov/golangForPro/phonebook/internal/models"
)

func main() {
	command, err := args.GetCommand()
	if err != nil {
		return
	}

	switch command.Name {
	case models.CommandSearch:
		key := command.Params[0]
		entry := commands.Search(key)
		if entry == nil {
			fmt.Println("No entry found: ", key)

			return
		}

		fmt.Println(entry.View())
	case models.CommandList:
		entries := commands.List()
		for _, entry := range entries {
			fmt.Println(entry.View())
		}
	}
}
