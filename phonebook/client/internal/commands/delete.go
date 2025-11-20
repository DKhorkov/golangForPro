package commands

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/client/internal/validation"
	"os"

	"github.com/spf13/cobra"
)

const (
	deleteKey = "key"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete an entry",
	Long:    `delete an entry from the phone book application.`,
	Aliases: []string{"d", "del"},
	Run: func(cmd *cobra.Command, args []string) {
		key, err := cmd.Flags().GetString(deleteKey)
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
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP(deleteKey, "k", "", "Key to delete")
}
