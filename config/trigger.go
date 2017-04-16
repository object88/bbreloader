package config

import "github.com/object88/bbreloader/glob"

type Trigger struct {
	Action  string
	Matcher *glob.Matcher
}

func parseTriggerConfig(config *config, trigger *triggerConfig) *Trigger {
	action := trigger.action
	m := glob.PreprocessGlobSpec(trigger.glob)
	return &Trigger{action, m}
}
