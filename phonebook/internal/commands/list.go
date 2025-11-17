package commands

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/internal/filepath"
	"github.com/DKhorkov/golangForPro/phonebook/internal/phonebook"
	"github.com/DKhorkov/golangForPro/phonebook/internal/readwriters"
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

		reverse, err := cmd.Flags().GetBool(reverseKey)
		if err != nil {
			fmt.Printf("Not a valid reverse argument: %v. Error: %v\n", reverse, err)

			os.Exit(1)
		}

		entries, err := pb.List(reverse)
		if err != nil {
			fmt.Println(err)

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
