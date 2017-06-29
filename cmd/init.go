package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const emptyConfig = "{\n  \"projects\": [{}]\n}\n"

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initializes a reloader configuration file.",
	Aliases: []string{"i", "initialize"},
	Long:    "Creates an empty '.reloader.json' file.",
	Run: func(cmd *cobra.Command, args []string) {
		// Attempt to open the file for writing.  We are specifying that the
		// file must be created and must not already exist.
		f, createErr := os.OpenFile(cfgFile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
		if os.IsExist(createErr) {
			fmt.Printf("Configuration file '%s' already exists; will not overwrite.\n", cfgFile)
			return
		}

		f.WriteString(emptyConfig)
		f.Close()

		fmt.Printf("Configuration file '%s' has been created.\n", cfgFile)
	},
}
