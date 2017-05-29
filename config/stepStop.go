package config

import (
	"context"
)

// StepStop is a custom step in the build process
type StepStop struct {
}

func newStepStop() *StepStop {
	return &StepStop{}
}

// Run will stop a launched instance
func (s *StepStop) Run(ctx context.Context, project *Project) (int, error) {
	project.Stop()
	return 0, nil
}
