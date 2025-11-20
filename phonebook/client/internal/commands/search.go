package commands

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/client/internal/validation"
	"os"

	"github.com/spf13/cobra"
)

const (
	searchKey = "key"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for the number",
	Long: `search whether a telephone number exists in the
	phone book application or not.`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		key, err := cmd.Flags().GetString(searchKey)
		if err != nil {
			fmt.Printf("Not a valid key: %s. Error: %v\n", key, err)

			os.Exit(1)
		}

		if !validation.ValidatePhone(key) {
			fmt.Printf("Not a valid Phone. Phone should be like \"+7 (911) 258-01-62\"\n")

			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP(searchKey, "k", "", "Key to search")
}
