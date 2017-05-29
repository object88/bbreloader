package config

import (
	"context"
)

// Step describes an increment in the reloader process
type Step interface {
	Run(ctx context.Context, project *Project) (int, error)
}

func stepConfigToStep(project *ProjectMapstructure, sc *StepMapstructure) Step {
	switch sc.Type {
	case "build":
		return newStepBuild(project, sc)
	case "stop":
		return newStepStop()
	default:
		return newStepCustom(sc)
	}
}
