package commands

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/client/internal/validation"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

const (
	deleteKey = "key"

	deleteURL = "http://%s:%d/entries/%s"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete an entry",
	Long:    `delete an entry from the phone book application.`,
	Aliases: []string{"d", "del"},
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

		key, err := cmd.Flags().GetString(deleteKey)
		if err != nil {
			fmt.Printf("Not a valid key: %s. Error: %v\n", key, err)

			os.Exit(1)
		}

		if !validation.ValidatePhone(key) {
			fmt.Printf("Not a valid Phone. Phone should be like \"+7 (911) 258-01-62\"\n")

			os.Exit(1)
		}

		addr := fmt.Sprintf(deleteURL, host, port, key)
		req, err := http.NewRequest(http.MethodDelete, addr, nil)
		if err != nil {
			fmt.Printf("Failed to create request: %v\n", err)

			os.Exit(1)
		}

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

		fmt.Printf("Successfully deleted entry: %s\n", key)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP(deleteKey, "k", "", "Key to delete")
}
