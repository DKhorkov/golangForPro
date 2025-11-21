package commands

import (
	"encoding/json"
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/client/internal/models"
	"github.com/DKhorkov/golangForPro/phonebook/client/internal/validation"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const (
	searchKey = "key"

	searchURL = "http://%s:%d/search"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for the number",
	Long: `search whether a telephone number exists in the
	phone book application or not.`,
	Aliases: []string{"s"},
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

		key, err := cmd.Flags().GetString(searchKey)
		if err != nil {
			fmt.Printf("Not a valid key: %s. Error: %v\n", key, err)

			os.Exit(1)
		}

		if !validation.ValidatePhone(key) {
			fmt.Printf("Not a valid Phone. Phone should be like \"+7 (911) 258-01-62\"\n")

			os.Exit(1)
		}

		addr := fmt.Sprintf(searchURL, host, port)
		req, err := http.NewRequest(methodGet, addr, nil)
		if err != nil {
			fmt.Printf("Failed to create request: %v\n", err)

			os.Exit(1)
		}

		q := req.URL.Query()
		q.Set(searchKey, key)
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
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP(searchKey, "k", "", "Key to search")
}
