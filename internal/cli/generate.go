package cli

import (
	"griffin/internal/log"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate Code",
	Long:  `Generate`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("generate code")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
