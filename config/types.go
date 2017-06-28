package config

import "bytes"

// ReloaderMapstructure is for internal use only
type ReloaderMapstructure struct {
	Projects []*ProjectMapstructure `mapstructure:"projects"`
}

func (r *ReloaderMapstructure) String() string {
	var buffer bytes.Buffer

	for k, v := range r.Projects {
		buffer.WriteString("Project #")
		buffer.WriteString(string(k))
		buffer.WriteString(":\n")
		buffer.WriteString(v.String())
	}

	return buffer.String()
}

// ProjectMapstructure is for internal use only
type ProjectMapstructure struct {
	Root   *string            `mapstructure:"root"`
	Target *string            `mapstructure:"target"`
	Build  *BuildMapstructure `mapstructure:"build"`
	Test   *TestMapstructure  `mapstructure:"test"`
	Run    *RunMapstructure   `mapstructure:"run"`
}

func (p *ProjectMapstructure) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Root: ")
	if p.Root == nil {
		buffer.WriteString("not provided")
	} else {
		buffer.WriteString(*p.Root)
	}
	buffer.WriteString("\n")

	return buffer.String()
}

// BuildMapstructure is for internal use only
type BuildMapstructure struct {
	Args           *[]string           `mapstructure:"args"`
	PreBuildSteps  *[]StepMapstructure `mapstructure:"pre-build-steps"`
	PostBuildSteps *[]StepMapstructure `mapstructure:"post-build-steps"`
}

// RunMapstructure is for internal use only
type RunMapstructure struct {
	Args        *[]string `mapstructure:"args"`
	Retain      *bool     `mapstructure:"retain"`
	RebuildGlob *string   `mapstructure:"rebuild-trigger-glob"`
	RestartGlob *string   `mapstructure:"restart-trigger-glob"`
}

// TestMapstructure is for internal use only
type TestMapstructure struct {
	Args        *[]string `mapstructure:"args"`
	RestartGlob *string   `mapstructure:"restart-trigger-glob"`
}

// StepMapstructure is for internal use only
type StepMapstructure struct {
	Command *string   `mapstructure:"command"`
	Args    *[]string `mapstructure:"args"`
}
