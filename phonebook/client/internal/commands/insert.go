package commands

import (
	"encoding/json"
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/client/internal/models"
	"github.com/DKhorkov/golangForPro/phonebook/client/internal/validation"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

const (
	nameKey    = "name"
	surnameKey = "surname"
	phoneKey   = "phone"

	insertURL = "http://%s:%d/insert"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:     "insert",
	Short:   "insert new data",
	Long:    `This command inserts new data into the phone book application.`,
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.Flags().GetString(hostKey)
		if err != nil {
			fmt.Printf("Failed to get host key from flags: %v\n", err)

			os.Exit(1)
		}

		port, err := cmd.Flags().GetInt(portKey)
		if err != nil {
			fmt.Printf("Failed to get port key from flags: %v\n", err)

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

		addr := fmt.Sprintf(insertURL, host, port)
		req, err := http.NewRequest(methodGet, addr, nil)
		if err != nil {
			fmt.Printf("Failed to create request: %v\n", err)

			os.Exit(1)
		}

		q := req.URL.Query()
		q.Set(nameKey, name)
		q.Set(surnameKey, surname)
		q.Set(phoneKey, phone)
		req.URL.RawQuery = q.Encode()

		httpClient := &http.Client{}
		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Printf("Failed to send request: %v\n", err)

			os.Exit(1)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			errInfo, _ := io.ReadAll(resp.Body)
			fmt.Printf("Failed to send request: %v. Error: %s\n", resp.Status, errInfo)

			os.Exit(1)
		}

		var entry models.Entry
		if err := json.NewDecoder(resp.Body).Decode(&entry); err != nil {
			fmt.Printf("Failed to decode response: %v\n", err)

			os.Exit(1)
		}

		fmt.Println(entry.View())
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	insertCmd.Flags().StringP(nameKey, "n", "", "name value")
	insertCmd.Flags().StringP(surnameKey, "s", "", "surname value")
	insertCmd.Flags().StringP(phoneKey, "p", "", "phone value")
}
