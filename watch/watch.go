package watch

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/object88/bbreloader/config"
	"github.com/rjeczalik/notify"
)

const (
	source string = "source"
)

func Run(projects *[]*config.Project) error {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	for _, project := range *projects {
		// Do initial build and start the run.

		build(project)
		project.Start()

		// Watch the files for changes
		err := Watch(project)
		if err != nil {
			return err
		}
	}

	// Wait for a signal to end the app.
	<-sigchan

	return nil
}

func watch(trigger *config.Trigger, notifyChan chan notify.EventInfo, callback func(*config.CollectedEvents)) {
	lull := time.Duration(2 * time.Second)

	for {
		select {
		case e := <-notifyChan:
			path := e.Path()
			log.Printf("%s :: %s\n", path, e.Event().String())
			if trigger.Matcher.Matches(path) {
				// We have a match!
				trigger.CollectedEvents.AddEvent(e)
			}
		case <-time.After(lull):
			processed := trigger.ResetTrigger()

			if processed.HasEvents() {
				go callback(processed)
			}
		}
	}
}

// Watch builds and starts the process, then watches the file system for
// changes to trigger another build or restart
func Watch(c *config.Project) error {
	notifyChan := make(chan notify.EventInfo, 4096)

	// Start watch at root filesystem level
	err := notify.Watch(c.Watch, notifyChan, notify.All)
	if err != nil {
		// Failed to start the watch; stop the channel and quit.
		notify.Stop(notifyChan)
		return err
	}

	go watch(c.Run.Rebuild, notifyChan, func(events *config.CollectedEvents) {
		build(c)
	})

	return nil
}

func build(project *config.Project) {
	// For cancelling log-running operations.
	ctx := context.Background()

	runSteps(ctx, project, &project.Build.PreBuild)

	project.Build.Run(ctx, project)

	runSteps(ctx, project, &project.Build.PostBuild)
}

func runSteps(ctx context.Context, project *config.Project, steps *[]*config.Step) {
	for i, step := range *steps {
		log.Printf("Step #%d...", i)
		step.Run(ctx, project)

		log.Printf("Finished step.\n")
	}
}

func restart(project *config.Project) {
	project.Stop()
	project.Start()
}
