package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	reverseKey = "reverse"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list all entries",
	Long:    `This command lists all entries in the phone book application.`,
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		reverse, err := cmd.Flags().GetBool(reverseKey)
		if err != nil {
			fmt.Printf("Not a valid reverse argument: %v. Error: %v\n", reverse, err)

			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP(reverseKey, "r", false, "Whether reverse order or not")
}
