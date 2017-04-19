package config

// ReloaderMapstructure is for internal use only
type ReloaderMapstructure struct {
	Projects []*ProjectMapstructure `mapstructure:"projects"`
}

// ProjectMapstructure is for internal use only
type ProjectMapstructure struct {
	Root     string                  `mapstructure:"root"`
	Target   string                  `mapstructure:"target"`
	Build    *BuildMapstructure      `mapstructure:"build"`
	Triggers *[]*TriggerMapstructure `mapstructure:"triggers"`
}

// BuildMapstructure is for internal use only
type BuildMapstructure struct {
	Steps []*StepMapstructure `mapstructure:"steps"`
}

// StepMapstructure is for internal use only
type StepMapstructure struct {
	Type    string    `mapstructure:"type"`
	Command *string   `mapstructure:"command"`
	Args    *[]string `mapstructure:"args"`
}

// TriggerMapstructure is for internal use only
type TriggerMapstructure struct {
	Action string `mapstructure:"action"`
	Glob   string `mapstructure:"glob"`
}
