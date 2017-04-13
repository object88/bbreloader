package config

// ReloaderConfig is a collection of Config
type ReloaderConfig struct {
	Configs []*Config `json:"configs"`
}

// Config captures the configuration for a particular target
type Config struct {
	Root     string           `json:"root"`
	Target   string           `json:"target"`
	Patterns []*PatternConfig `json:"patterns"`
}

// PatternConfig associates a blag pattern with a sequence of steps
type PatternConfig struct {
	Glob  string        `json:"glob"`
	Steps []*StepConfig `json:"steps"`
}

// StepConfig is a single step in a build
type StepConfig struct {
	Type    string    `json:"type"`
	Command *string   `json:"command"`
	Args    *[]string `json:"args"`
	Retain  *bool     `json:"retain"`
}
