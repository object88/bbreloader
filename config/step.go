package config

import (
	"context"
)

// Step describes an increment in the reloader process
type Step interface {
	Run(ctx context.Context) (int, error)
}

func stepConfigToStep(config *config, sc *stepConfig) Step {
	switch sc.t {
	case "build":
		return newStepBuild(config, sc)
	default:
		return newStepCustom()
	}
}
