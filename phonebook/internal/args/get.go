package args

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/internal/commands"
	"github.com/DKhorkov/golangForPro/phonebook/internal/models"
	"github.com/DKhorkov/golangForPro/phonebook/internal/validation"
	"os"
	"path"
)

func ReadCommand() (*models.Command, error) {
	exe := path.Base(os.Args[0])

	if len(os.Args) < 2 {
		return nil, fmt.Errorf("%w. Usage: %s search|list|insert|delete <argument>", ErrInvalidUsage, exe)
	}

	switch commands.CommandName(os.Args[1]) {
	case commands.CommandList:
		return &models.Command{
			Name: commands.CommandList,
		}, nil
	case commands.CommandSearch:
		if len(os.Args) < 3 {
			return nil, fmt.Errorf("%w. Usage: %s search Phone", ErrInvalidUsage, exe)
		}

		return &models.Command{
			Name:   commands.CommandSearch,
			Params: os.Args[2:],
		}, nil
	case commands.CommandInsert:
		if len(os.Args) < 5 {
			return nil, fmt.Errorf("%w. Usage: %s insert Name Surname Phone", ErrInvalidUsage, exe)
		}

		if !validation.ValidateNameSurname(os.Args[2]) {
			return nil,
				fmt.Errorf(
					"%w. Not a valid Name. Name should be like \"Илья\" or \"Ilya\"",
					ErrInvalidArguments,
				)
		}

		if !validation.ValidateNameSurname(os.Args[3]) {
			return nil,
				fmt.Errorf(
					"%w. Not a valid Surname. Surname should be like \"Романов\" or \"Romanov\"",
					ErrInvalidArguments,
				)
		}

		if !validation.ValidatePhone(os.Args[4]) {
			return nil,
				fmt.Errorf(
					"%w. Not a valid Phone. Phone should be like \"+7 (911) 258-01-62\"",
					ErrInvalidArguments,
				)
		}

		return &models.Command{
			Name:   commands.CommandInsert,
			Params: os.Args[2:],
		}, nil
	case commands.CommandDelete:
		if len(os.Args) < 3 {
			return nil, fmt.Errorf("%w. Usage: %s delete Phone", ErrInvalidUsage, exe)
		}

		return &models.Command{
			Name:   commands.CommandDelete,
			Params: os.Args[2:],
		}, nil
	}

	return nil, fmt.Errorf("%w. Usage: %s search|list|insert|delete <argument>", ErrInvalidCommand, exe)
}
