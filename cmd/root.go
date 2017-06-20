package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

func init() {
	RootCmd.AddCommand(initCmd, runCmd, testCmd, versionCmd)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.reloader.json)")
}

// RootCmd is the main action taken by Cobra
var RootCmd = &cobra.Command{
	Use:   "bbreloader",
	Short: "bbreloader is a watcher for developers",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}
