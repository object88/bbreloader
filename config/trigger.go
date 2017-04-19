package config

import "github.com/object88/bbreloader/glob"

// Trigger is a glob of files to watch and an action to take when the change
type Trigger struct {
	Action  string
	Matcher *glob.Matcher
}

func parseTriggerConfig(trigger *TriggerMapstructure) *Trigger {
	action := trigger.Action
	m := glob.PreprocessGlobSpec(trigger.Glob)
	return &Trigger{action, m}
}
