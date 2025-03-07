package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Database Migration",
	Long:  `Database Migration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run database migrations")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
