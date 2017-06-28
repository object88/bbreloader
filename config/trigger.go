package config

import "github.com/object88/bbreloader/glob"

// Trigger is a glob of files to watch and an action to take when the change
type Trigger struct {
	CollectedEvents *CollectedEvents
	Matcher         *glob.Matcher
}

func parseGlob(path string) *Trigger {
	if path == "" {
		return nil
	}
	ce := newCollectedEvents()
	m := glob.PreprocessGlobSpec(path)
	return &Trigger{ce, m}
}

// ResetTrigger returns all the events associated with this trigger and starts
// a new collection
func (t *Trigger) ResetTrigger() *CollectedEvents {
	e := t.CollectedEvents
	t.CollectedEvents = newCollectedEvents()
	return e
}
