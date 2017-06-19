package cmd

import "github.com/spf13/cobra"

var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   "Runs `go test` in the Root directory",
	Long:    "Executes the test files",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
