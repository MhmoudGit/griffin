package cli

import (
	"griffin/internal/generators"
	"griffin/internal/log"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		err := generators.NewProject(projectName)
		if err != nil {
			log.Error("Error:", err)
			return
		}
		log.Success("Project created successfully: ", projectName)
		log.Info("$ cd ", projectName, " // to start developing")
		log.Info("run a database with the project name")
		log.Info("$ griffin server // to start the server")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
