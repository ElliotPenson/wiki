package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Serve and re-render HTML on file changes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("watch called")
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
