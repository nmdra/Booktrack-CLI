/*
Copyright © 2025 NIMENDRA <nimendraonline@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Booktrack-CLI",
	Short: "CLI for managing books using sqlc + sqlite",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
