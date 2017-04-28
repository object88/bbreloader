package config

import (
	"context"
	"log"
)

// StepCustom is a custom step in the build process
type StepCustom struct {
	Command string    `json:"command"`
	Args    *[]string `json:"args"`
	Retain  *bool     `json:"retain"`
}

func newStepCustom() *StepCustom {
	return &StepCustom{}
}

// Run is not implemented
func (s *StepCustom) Run(ctx context.Context) (int, error) {
	log.Printf("Custom step not implemented.")
	return 0, nil
}
