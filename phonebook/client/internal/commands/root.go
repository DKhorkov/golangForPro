package commands

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	hostKey = "host"
	portKey = "port"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "phonebook",
	Short: "A phone book application",
	Long:  `This is a Phone Book application that uses JSON records.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP(hostKey, "H", "localhost", "Host")
	rootCmd.PersistentFlags().IntP(portKey, "P", 8080, "Port number")
}
