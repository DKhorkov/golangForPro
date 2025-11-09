package args

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/internal/models"
	"os"
	"path"
)

func GetCommand() (*models.Command, error) {
	exe := path.Base(os.Args[0])

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s search|list <argument>\n", exe)

		return nil, ErrInvalidArguments
	}

	switch models.CommandName(os.Args[1]) {
	case models.CommandSearch:
		if len(os.Args) < 3 {
			fmt.Printf("Usage: %s search <argument>\n", exe)

			return nil, ErrInvalidArguments
		}

		return &models.Command{
			Name:   models.CommandSearch,
			Params: os.Args[2:],
		}, nil
	case models.CommandList:
		return &models.Command{
			Name: models.CommandList,
		}, nil
	}

	fmt.Printf("Usage: %s search|list <argument>\n", exe)
	return nil, ErrInvalidCommand
}
