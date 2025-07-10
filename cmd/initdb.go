/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/nmdra/Booktrack-CLI/internal/db"
	"github.com/spf13/cobra"
)

var initdbCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database if not present",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.InitDatabase(); err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
}
