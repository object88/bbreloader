package watch

import (
	"log"
	"time"

	"github.com/object88/bbreloader/config"
	"github.com/rjeczalik/notify"
)

// Watch builds and starts the process, then watches the file system for
// changes to trigger another build or restart
func Watch(p *config.Project, trigger *config.Trigger, callback func(events *config.CollectedEvents)) error {
	notifyChan := make(chan notify.EventInfo, 4096)

	// Start watch at root filesystem level
	err := notify.Watch(p.Watch, notifyChan, notify.All)
	if err != nil {
		// Failed to start the watch; stop the channel and quit.
		notify.Stop(notifyChan)
		return err
	}

	go watch(p, trigger, notifyChan, callback)

	return nil
}

func watch(project *config.Project, trigger *config.Trigger, notifyChan chan notify.EventInfo, callback func(*config.CollectedEvents)) {
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
