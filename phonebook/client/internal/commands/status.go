package commands

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const (
	methodGet = "GET"

	statusURL = "http://%s:%d/status"
)

// listCmd represents the list command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Prints the status of the server.",
	Long: `This command prints information about the
	status of the phone book server.`,
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

		addr := fmt.Sprintf(statusURL, host, port)
		req, err := http.NewRequest(methodGet, addr, nil)
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

		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Current number of entries: %s\n", string(body))
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
