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
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
