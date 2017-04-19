package config

// Build contains a series of individual steps necessary to build a project
type Build struct {
	Steps []Step
}

func parseBuildConfig(project *ProjectMapstructure, build *BuildMapstructure) *Build {
	count := 0
	if build.Steps != nil {
		count = len(build.Steps)
	}
	s := make([]Step, count)
	for i := 0; i < count; i++ {
		s[i] = stepConfigToStep(project, build.Steps[i])
	}
	return &Build{s}
}
