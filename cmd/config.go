package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

// Read in config file and ENV variables if set.
func readConfig() {
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
