package commands

import (
	"encoding/json"
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/client/internal/models"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

const (
	reverseKey = "reverse"

	listURL = "http://%s:%d/entries"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all entries",
	Long:    `This command lists all entries in the phone book application.`,
	Aliases: []string{"l"},
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

		reverse, err := cmd.Flags().GetBool(reverseKey)
		if err != nil {
			fmt.Printf("Not a valid reverse argument: %v. Error: %v\n", reverse, err)

			os.Exit(1)
		}

		addr := fmt.Sprintf(listURL, host, port)
		req, err := http.NewRequest(http.MethodGet, addr, nil)
		if err != nil {
			fmt.Printf("Failed to create request: %v\n", err)

			os.Exit(1)
		}

		q := req.URL.Query()
		q.Set(reverseKey, strconv.FormatBool(reverse))
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

		var entries []models.Entry
		if err := json.NewDecoder(resp.Body).Decode(&entries); err != nil {
			fmt.Printf("Failed to decode response: %v\n", err)

			os.Exit(1)
		}

		for _, entry := range entries {
			fmt.Println(entry.View())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP(reverseKey, "r", false, "Whether reverse order or not")
}
