package watch

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/object88/bbreloader/config"
	"github.com/rjeczalik/notify"
)

func Run(projects *[]*config.Project) error {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// Initialize the build directory
	config.InitializeBuildDirectory()

	for _, project := range *projects {
		// Do initial build and start the run.

		project.Build.Run(project)
		project.Start()

		// Watch the files for changes
		err := Watch(project)
		if err != nil {
			return err
		}
	}

	// Wait for a signal to end the app.
	<-sigchan

	// Clean up the build temp directory
	config.DestroyBuildDirectory()

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
func Watch(p *config.Project) error {
	notifyChan := make(chan notify.EventInfo, 4096)

	// Start watch at root filesystem level
	err := notify.Watch(p.Watch, notifyChan, notify.All)
	if err != nil {
		// Failed to start the watch; stop the channel and quit.
		notify.Stop(notifyChan)
		return err
	}

	go watch(p.Run.Rebuild, notifyChan, func(events *config.CollectedEvents) {
		p.Build.Run(p)
	})

	return nil
}

func restart(project *config.Project) {
	project.Stop()
	project.Start()
}
