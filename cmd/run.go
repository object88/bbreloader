package cmd

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r"},
	Short:   "Runs the application",
	Long:    "For projects with a Target specified, will start the application and restart as code changes.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
