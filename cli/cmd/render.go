package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Convert wiki markdown into HTML",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("render called")
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)
}
