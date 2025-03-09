package cli

import (
	"github.com/MhmoudGit/griffin/internal/log"
	"github.com/MhmoudGit/griffin/internal/server"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run Server",
	Long:  `Run Server`,
	Run: func(cmd *cobra.Command, args []string) {
		err := server.Start()
		if err != nil {
			log.Error("Error:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
