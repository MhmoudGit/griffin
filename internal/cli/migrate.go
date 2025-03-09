package cli

import (
	"database/sql"
	"fmt"
	"griffin/internal/config"
	"griffin/internal/log"

	"github.com/pressly/goose"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Database Migration",
	Long:  `Run database migrations using Goose.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Running database migrations...")

		// Load the configuration
		config, err := config.LoadConfig()
		if err != nil {
			log.Error("Failed to load config: ", err)
			return
		}

		// Build the DSN (Database Connection String)
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			config.Database.Host,
			config.Database.User,
			config.Database.Password,
			config.Database.Name,
			config.Database.Port,
			config.Database.SSLMode,
		)

		// Open a database connection
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			log.Error("Failed to open database connection: ", err)
			return
		}
		defer db.Close() // Ensure the connection is closed after the function exits

		// Run migrations using Goose
		if err := goose.Up(db, "migrations"); err != nil {
			log.Error("Failed to run migrations: ", err)
			return
		}

		log.Success("Migrations completed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
