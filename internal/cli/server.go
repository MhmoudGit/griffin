package cli

import (
	"griffin/internal/log"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run Server",
	Long:  `Run Server`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("run server")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
