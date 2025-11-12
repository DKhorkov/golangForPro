package main

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/internal/args"
	"github.com/DKhorkov/golangForPro/phonebook/internal/phonebook"
	"github.com/DKhorkov/golangForPro/phonebook/internal/readwriters"
)

const (
	csvFilepath = "phonebook.csv"
)

func main() {
	command, err := args.ReadCommand()
	if err != nil {
		fmt.Println(err)

		return
	}

	rw := readwriters.NewCSVReadWriter(csvFilepath)
	pb, err := phonebook.New(rw)
	if err != nil {
		fmt.Println(err)

		return
	}

	if err = pb.Execute(*command); err != nil {
		fmt.Println(err)
	}
}
