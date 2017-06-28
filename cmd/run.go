package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"

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

		projects, ok := config.SetupProjects()
		if !ok {
			fmt.Printf("NOPE.")
			return
		}

		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)

		for _, project := range *projects {
			// Do initial build and start the run.
			p := project

			// Initialize the build directory
			p.Build.InitializeBuildDirectory()

			p.Build.Run(p)
			p.Start()

			// Watch the files for changes
			rebuildErr := watch.Watch(p, p.Run.Rebuild, func(collectedEvents *config.CollectedEvents) {
				p.Build.Run(p)
			})
			if rebuildErr != nil {
				log.Printf("Failed to start 'rebuild' watch; %s\n", rebuildErr.Error())
			}

			restartErr := watch.Watch(p, p.Run.Restart, func(collectedEvents *config.CollectedEvents) {
				p.Stop()
				p.Start()
			})
			if restartErr != nil {
				log.Printf("Failed to start 'restart' watch; %s\n", restartErr.Error())
			}
		}

		// Wait for a signal to end the app.
		<-sigchan

		for _, project := range *projects {
			// Clean up the build temp directory
			project.Build.DestroyBuildDirectory()
		}
	},
}
