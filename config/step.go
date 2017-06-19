package config

import (
	"context"
	"log"
	"os"
	"os/exec"
)

// Step describes an increment in the reloader process
type Step struct {
	Command string    `json:"command"`
	Args    *[]string `json:"args"`
}

func stepConfigToStep(project *ProjectMapstructure, sc *StepMapstructure) *Step {
	return &Step{
		Command: *sc.Command,
		Args:    sc.Args,
	}
}

// Run is not implemented
func (s *Step) Run(ctx context.Context, _ *Project) (int, error) {
	// For now, just route output to stdout.
	args := parseArgs(s.Args)
	cmd := exec.CommandContext(ctx, s.Command, *args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Build command failed: %s\n", err.Error())
		return 0, err
	}

	return 0, nil
}
