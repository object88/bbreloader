package watch

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/object88/bbreloader/config"
	"github.com/object88/bbreloader/glob"
	"github.com/rjeczalik/notify"
)

const (
	source string = "source"
)

func Run(configs *[]*config.Project) error {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	for _, config := range *configs {
		err := Watch(config)
		if err != nil {
			return err
		}
	}

	// Wait for a signal to end the app.
	<-sigchan

	return nil
}

func watch(globs *glob.Matcher, notifyChan chan notify.EventInfo, callback func(*collectedEvents)) {
	events := newCollectedEvents()
	lull := time.Duration(2 * time.Second)

	for {
		select {
		case e := <-notifyChan:
			path := e.Path()
			log.Printf("File '%s' changed!\n", path)
			if globs.Matches(path) {
				log.Printf("Got match\n")
				// We have a match!
				switch e.Event() {
				case notify.Create:
					events.created[path] = true
				case notify.Remove:
					events.removed[path] = true
				case notify.Rename:
					events.renamed[path] = true
				case notify.Write:
					events.written[path] = true
				}
			}
		case <-time.After(lull):
			processed := events
			events = newCollectedEvents()
			go callback(processed)
		}
	}
}

func Watch(config *config.Project) error {
	notifyChan := make(chan notify.EventInfo, 4096)

	// Start watch at root filesystem level
	err := notify.Watch(config.Watch, notifyChan, notify.All)
	if err != nil {
		// Failed to start the watch; stop the channel and quit.
		notify.Stop(notifyChan)
		return err
	}

	// Loop.
	for _, v := range config.Triggers {
		go watch(v.Matcher, notifyChan, func(events *collectedEvents) {
			if !events.HasEvents() {
				return
			}
			fmt.Printf("Changed files: %#v\n", events.written)
			execute(config)
		})
	}

	return nil
}

func execute(config *config.Project) {
	// For cancelling log-running operations.
	ctx := context.Background()
	steps := config.Build.Steps

	for _, step := range steps {
		step.Run(ctx)

		log.Printf("Finished step.\n")
	}
}
