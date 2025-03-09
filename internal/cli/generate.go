package cli

import (
	"fmt"
	"griffin/internal/generators"
	"griffin/internal/log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "gen [resource] [field:type...]",
	Short: "Generate CRUD code for a resource",
	Long:  `Generate CRUD code for a resource, including model, handler, repository, and routes.`,
	Args:  cobra.MinimumNArgs(1), // At least the resource name is required
	Run: func(cmd *cobra.Command, args []string) {
		resource := args[0]
		fields := args[1:]
		dir, _ := os.Getwd()
		projectName := filepath.Base(dir)

		// Generate a migration file for the resource
		if err := generators.GenerateMigration(resource, fields); err != nil {
			log.Error("Failed to generate migration: ", err)
			return
		}

		resourceDir := strings.ToLower(resource)
		if err := os.Mkdir(resourceDir, os.ModePerm); err != nil {
			log.Error("Failed to create resource: ", err)
			return
		}

		if err := os.Chdir(resourceDir); err != nil {
			log.Error("Failed to change to resource directory: ", err)
			return
		}

		log.Info("Generating CRUD code for resource: ", resource)
		log.Info("Fields: ", fields)

		// Generate model, handler, repository, and routes
		if err := generators.GenerateModel(resource, fields); err != nil {
			log.Error("Failed to generate model: ", err)
			return
		}
		if err := generators.GenerateHandlers(resource, projectName); err != nil {
			log.Error("Failed to generate handler: ", err)
			return
		}
		if err := generators.GenerateRepository(resource); err != nil {
			log.Error("Failed to generate repository: ", err)
			return
		}
		if err := generators.GenerateRoutes(resource); err != nil {
			log.Error("Failed to add routes to main.go: ", err)
			return
		}

		log.Success("CRUD code generated successfully.")
		log.Info(fmt.Sprintf(`
		add this to main.go under //Routes
			%s.SetupRoutes(e, pg.DB)

		Then: run griffin server
		`, resource))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
