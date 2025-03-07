package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "New Project",
	Long:  `New Project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create new project")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
