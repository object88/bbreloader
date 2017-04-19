package config

import "context"

// StepCustom is a custom step in the build process
type StepCustom struct {
	Command string    `json:"command"`
	Args    *[]string `json:"args"`
	Retain  *bool     `json:"retain"`
}

func newStepCustom() *StepCustom {
	return &StepCustom{}
}

func (s *StepCustom) Run(ctx context.Context) (int, error) {
	return 0, nil
}
