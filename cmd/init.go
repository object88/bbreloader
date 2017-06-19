package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

const emptyConfig = "{\n  projects: [{}]\n}\n"

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initializes a reloader configuration file.",
	Aliases: []string{"i", "initialize"},
	Long:    "Creates an empty '.reloader.json' file.",
	Run: func(cmd *cobra.Command, args []string) {
		// Check to see if the configuration file already exists;
		// we don't want to overwrite the file.
		if false {
			fmt.Printf("Configuration file '%s' already exists; will not overwrite.\n", ".reloaded.json")
			return
		}

		// Open `.reloader.json` and write the contents.
		ioutil.WriteFile(".reloader.json", []byte(emptyConfig), 0)
	},
}
