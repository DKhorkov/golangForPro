package commands

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/filepath"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/models"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/phonebook"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/readwriters"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/validation"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const (
	nameKey    = "name"
	surnameKey = "surname"
	phoneKey   = "phone"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:     "insert",
	Short:   "insert new data",
	Long:    `This command inserts new data into the phone book application.`,
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		fp, err := filepath.Get()
		if err != nil {
			fmt.Println(err)

			os.Exit(1)
		}

		rw, err := readwriters.New(fp)
		if err != nil {
			fmt.Println(err)

			os.Exit(1)
		}

		pb, err := phonebook.New(rw)
		if err != nil {
			fmt.Println(err)

			os.Exit(1)
		}

		name, err := cmd.Flags().GetString(nameKey)
		if err != nil {
			fmt.Printf("Not a valid name: %s. Error: %v\n", name, err)

			os.Exit(1)
		}

		if !validation.ValidateNameSurname(name) {
			fmt.Printf("Not a valid Name. Name should be like \"Илья\" or \"Ilya\"\n")

			os.Exit(1)
		}

		surname, err := cmd.Flags().GetString(surnameKey)
		if err != nil {
			fmt.Printf("Not a valid surname: %s. Error: %v\n", surname, err)

			os.Exit(1)
		}

		if !validation.ValidateNameSurname(surname) {
			fmt.Printf("Not a valid Surname. Surname should be like \"Романов\" or \"Romanov\"\n")

			os.Exit(1)
		}

		phone, err := cmd.Flags().GetString(phoneKey)
		if err != nil {
			fmt.Printf("Not a valid phone: %s. Error: %v\n", phone, err)

			os.Exit(1)
		}

		if !validation.ValidatePhone(phone) {
			fmt.Printf("Not a valid Phone. Phone should be like \"+7 (911) 258-01-62\"\n")

			os.Exit(1)
		}

		entry := models.Entry{
			Name:       name,
			Surname:    surname,
			Phone:      phone,
			LastAccess: time.Now(),
		}

		if err := pb.Insert(entry); err != nil {
			fmt.Println(err)

			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	insertCmd.Flags().StringP(nameKey, "n", "", "name value")
	insertCmd.Flags().StringP(surnameKey, "s", "", "surname value")
	insertCmd.Flags().StringP(phoneKey, "p", "", "phone value")
}
