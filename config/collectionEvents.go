package config

import "github.com/rjeczalik/notify"

// CollectedEvents is all the changes to any file in the file system Starting
// from the Target
type CollectedEvents struct {
	created map[string]bool
	removed map[string]bool
	renamed map[string]bool
	written map[string]bool
}

func newCollectedEvents() *CollectedEvents {
	return &CollectedEvents{
		created: make(map[string]bool),
		removed: make(map[string]bool),
		renamed: make(map[string]bool),
		written: make(map[string]bool),
	}
}

// AddEvent adds an event to the collection
func (ce *CollectedEvents) AddEvent(e notify.EventInfo) {
	path := e.Path()
	switch e.Event() {
	case notify.Create:
		ce.created[path] = true
	case notify.Remove:
		ce.removed[path] = true
	case notify.Rename:
		ce.renamed[path] = true
	case notify.Write:
		ce.written[path] = true
	}
}

// HasEvents determines whether there are any changes
func (ce *CollectedEvents) HasEvents() bool {
	ce.minimize()

	return len(ce.created) > 0 ||
		len(ce.removed) > 0 ||
		len(ce.renamed) > 0 ||
		len(ce.written) > 0
}

func (ce *CollectedEvents) String() string {
	return ""
}

func (ce *CollectedEvents) minimize() {

}
