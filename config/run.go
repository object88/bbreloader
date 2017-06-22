package config

import (
	"context"
	"log"
	"os"
	"os/exec"
)

const defaultRestartGlob = ""
const defaultRebuildGlob = "*.go"

type Run struct {
	Args    *Args
	Retain  bool
	Rebuild *Trigger
	Restart *Trigger
}

func parseRun(project *ProjectMapstructure, r *RunMapstructure) *Run {
	args := parseArgs(r.Args)

	retain := false
	if r.Retain != nil {
		retain = *r.Retain
	}

	rebuildGlob := defaultRebuildGlob
	if r.RebuildGlob != nil {
		rebuildGlob = *r.RebuildGlob
	}
	rebuild := parseGlob(rebuildGlob)

	restartGlob := defaultRestartGlob
	if r.RestartGlob != nil {
		restartGlob = *r.RestartGlob
	}
	restart := parseGlob(restartGlob)

	return &Run{args, retain, rebuild, restart}
}

// Run executes the step with an interruptable context
func (r *Run) Run(ctx context.Context, p *Project) (int, error) {
	// For now, just route output to stdout.
	cmd := exec.CommandContext(ctx, *p.Target, *r.Args...)
	cmd.Dir = p.Root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Run command failed: %s\n", err.Error())
		return 0, err
	}

	return 0, nil
}
