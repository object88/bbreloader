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
			p.Builder.InitializeBuildDirectory()

			p.Builder.Run(p)
			p.Runner.Start(p)

			// Watch the files for changes
			if p.Runner.Rebuild != nil {
				rebuildErr := watch.Watch(p, p.Runner.Rebuild, func(collectedEvents *config.CollectedEvents) {
					p.Builder.Run(p)
					p.Runner.Start(p)
				})
				if rebuildErr != nil {
					log.Printf("Failed to start 'rebuild' watch; %s\n", rebuildErr.Error())
				}
			}

			if p.Runner.Restart != nil {
				restartErr := watch.Watch(p, p.Runner.Restart, func(collectedEvents *config.CollectedEvents) {
					p.Runner.Stop()
					p.Runner.Start(p)
				})
				if restartErr != nil {
					log.Printf("Failed to start 'restart' watch; %s\n", restartErr.Error())
				}
			}
		}

		// Wait for a signal to end the app.
		<-sigchan

		for _, project := range *projects {
			// Clean up the build temp directory
			project.Builder.DestroyBuildDirectory()
		}
	},
}
