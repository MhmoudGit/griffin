package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run Server",
	Long:  `Run Server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run server")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
