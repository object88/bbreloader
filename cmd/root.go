package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	RootCmd.AddCommand(initCmd, runCmd, testCmd, versionCmd)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.reloader.json)")

	initConfig()
}

// Read in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".reloader") // name of config file (without extension)
	viper.AddConfigPath(".")         // adding home directory as first search path
	viper.AutomaticEnv()             // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to read config: %s\n", err.Error())
		return
	}
	fmt.Printf("Using config file: %#v\n", viper.ConfigFileUsed())

	// Make custom implementation with notify.
	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Project file changed:", e.Name)
	// })
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
