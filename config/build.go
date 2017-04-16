package config

type Build struct {
	Steps []Step
}

func parseBuildConfig(config *config, steps []*stepConfig) *Build {
	s := make([]Step, len(steps))
	for i := 0; i < len(s); i++ {
		s[i] = stepConfigToStep(config, steps[i])
	}
	return &Build{s}
}
