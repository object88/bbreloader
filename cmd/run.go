package cmd

import (
	"fmt"

	"github.com/object88/bbreloader/config"
	"github.com/object88/bbreloader/watch"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r"},
	Short:   "Runs the application",
	Long:    "For projects with a Target specified, will start the application and restart as code changes.",
	Run: func(cmd *cobra.Command, args []string) {
		readConfig()

		configs, ok := config.SetupConfig()
		if !ok {
			fmt.Printf("NOPE.")
			return
		}

		watch.Run(configs)
	},
}
