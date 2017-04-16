package config

import (
	"context"
	"log"
	"os"
	"os/exec"
)

// StepBuild describes a build step
type StepBuild struct {
	root   string
	target string
	args   []string
}

func newStepBuild(config *config, sc *stepConfig) *StepBuild {
	args := []string{}
	if sc.args != nil {
		args = *sc.args
	}
	return &StepBuild{config.root, config.target, args}
}

// Run executes the step with an interruptable context
func (s *StepBuild) Run(ctx context.Context) (int, error) {
	n := len(s.args)
	completeArgs := make([]string, n+3)
	completeArgs[0] = "build"
	for i, a := range s.args {
		completeArgs[i+1] = a
	}

	completeArgs[n+1] = "-o"
	completeArgs[n+2] = s.target

	// For now, just route output to stdout.
	cmd := exec.CommandContext(ctx, "go", completeArgs...)
	cmd.Dir = s.root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Build command failed: %s\n", err.Error())
		return 0, err
	}
	return 0, nil
}
