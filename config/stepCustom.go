package config

import (
	"context"
	"log"
	"os"
	"os/exec"
)

// StepCustom is a custom step in the build process
type StepCustom struct {
	Command string    `json:"command"`
	Args    *[]string `json:"args"`
	Retain  *bool     `json:"retain"`
}

func newStepCustom(sc *StepMapstructure) *StepCustom {
	return &StepCustom{
		Command: *sc.Command,
		Args:    sc.Args,
		Retain:  nil,
	}
}

// Run is not implemented
func (s *StepCustom) Run(ctx context.Context, _ *Project) (int, error) {
	// For now, just route output to stdout.
	args := []string{}
	if s.Args != nil {
		args = *s.Args
	}
	cmd := exec.CommandContext(ctx, s.Command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Build command failed: %s\n", err.Error())
		return 0, err
	}

	return 0, nil
}
