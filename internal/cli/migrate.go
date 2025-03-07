package cli

import (
	"griffin/internal/log"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Database Migration",
	Long:  `Database Migration`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("run database migrations")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
