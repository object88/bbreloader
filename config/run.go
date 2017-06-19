package config

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const defaultRestartGlob = ""
const defaultRebuildGlob = "*.go"

type Run struct {
	Args        *Args
	Retain      bool
	RebuildGlob string
	RestartGlob string
}

func newRun(project *ProjectMapstructure, r *RunMapstructure) *Run {
	args := parseArgs(r.Args)

	retain := false
	if r.Retain != nil {
		retain = *r.Retain
	}

	rebuildGlob := defaultRebuildGlob
	if r.RebuildGlob != nil {
		rebuildGlob = *r.RebuildGlob
	}

	restartGlob := defaultRestartGlob
	if r.RestartGlob != nil {
		restartGlob = *r.RestartGlob
	}

	return &Run{args, retain, rebuildGlob, restartGlob}
}

// Run executes the step with an interruptable context
func (r *Run) Run(ctx context.Context, p *Project) (int, error) {
	var tempFileName string
	copy := false
	tempFile, err := ioutil.TempFile("", "tmp")
	if err != nil {
		tempFileName = *p.Target
	} else {
		copy = true
		tempFileName = tempFile.Name()
	}

	completeArgs := r.Args.prependArgs("build", "-o", tempFileName)

	// For now, just route output to stdout.
	cmd := exec.CommandContext(ctx, "go", *completeArgs...)
	cmd.Dir = p.Root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Build command failed: %s\n", err.Error())
		return 0, err
	}

	if copy {
		linkErr := os.Link(tempFileName, *p.Target)
		if linkErr != nil {
			// Crap.
		}
	}

	return 0, nil
}
