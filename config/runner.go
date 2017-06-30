package config

import (
	"context"
	"log"
	"os/exec"
)

const defaultRestartGlob = ""
const defaultRebuildGlob = "*.go"

// Runner represents a runnable process
type Runner struct {
	Args     *Args
	Command  *string
	Retain   bool
	Rebuild  *Trigger
	Restart  *Trigger
	ctx      *context.Context
	cancelFn *context.CancelFunc
}

func parseRun(project *ProjectMapstructure, r *RunMapstructure) *Runner {
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

	return &Runner{args, r.Command, retain, rebuild, restart, nil, nil}
}

// Start spins up the process
func (r *Runner) Start(p *Project) {
	command := r.Command
	if command == nil {
		command = p.Target
	}
	if command == nil {
		// Nothing to run?
		return
	}

	if r.Retain {
		ctx, cancelFn := context.WithCancel(context.Background())
		log.Printf("Starting process...")
		startErr := exec.CommandContext(ctx, *command).Start()
		if startErr != nil {
			log.Printf("Failed to start retained process.")
			cancelFn()
			return
		}
		r.ctx = &ctx
		r.cancelFn = &cancelFn
	} else {
		startErr := exec.Command(*command).Start()
		if startErr != nil {
			log.Printf("Failed to start process.")
			return
		}
	}
	log.Printf("Started.")
}

// Stop shuts down the process
func (r *Runner) Stop() {
	if r.ctx != nil {
		log.Printf("Stopping process...")
		(*r.cancelFn)()
		r.ctx = nil
		r.cancelFn = nil
		log.Printf("Stopped.")
	}
}
