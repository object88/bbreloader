package config

// ReloaderMapstructure is for internal use only
type ReloaderMapstructure struct {
	Projects []*ProjectMapstructure `mapstructure:"projects"`
}

// ProjectMapstructure is for internal use only
type ProjectMapstructure struct {
	Root   *string            `mapstructure:"root"`
	Target *string            `mapstructure:"target"`
	Build  *BuildMapstructure `mapstructure:"build"`
	Test   *TestMapstructure  `mapstructure:"test"`
	Run    *RunMapstructure   `mapstructure:"run"`
}

// BuildMapstructure is for internal use only
type BuildMapstructure struct {
	Args           *[]string           `mapstructure:"args"`
	PreBuildSteps  *[]StepMapstructure `mapstructure:"pre-build-steps"`
	PostBuildSteps *[]StepMapstructure `mapstructure:"post-build-steps"`
}

type RunMapstructure struct {
	Args        *[]string `mapstructure:"args"`
	Retain      *bool     `mapstructure:"retain"`
	RebuildGlob *string   `mapstructure:"rebuild-trigger-glob"`
	RestartGlob *string   `mapstructure:"restart-trigger-glob"`
}

type TestMapstructure struct {
	Args        *[]string `mapstructure:"args"`
	RestartGlob *string   `mapstructure:"restart-trigger-glob"`
}

// StepMapstructure is for internal use only
type StepMapstructure struct {
	Command *string   `mapstructure:"command"`
	Args    *[]string `mapstructure:"args"`
}
