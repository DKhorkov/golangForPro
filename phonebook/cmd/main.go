package main

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/internal/args"
	"github.com/DKhorkov/golangForPro/phonebook/internal/phonebook"
	"github.com/DKhorkov/golangForPro/phonebook/internal/readwriters"
	"os"
)

const (
	csvEnv             = "PHONEBOOK_CSV"
	csvDefaultFilepath = "phonebook.csv"
)

func main() {
	filepath := os.Getenv(csvEnv)
	if filepath == "" {
		filepath = csvDefaultFilepath
	}

	command, err := args.ReadCommand()
	if err != nil {
		fmt.Println(err)

		return
	}

	rw := readwriters.NewCSVReadWriter(filepath)
	pb, err := phonebook.New(rw)
	if err != nil {
		fmt.Println(err)

		return
	}

	if err = pb.Execute(*command); err != nil {
		fmt.Println(err)
	}
}
