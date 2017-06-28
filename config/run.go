package config

import (
	"context"
	"log"
	"os/exec"
)

const defaultRestartGlob = ""
const defaultRebuildGlob = "*.go"

type Run struct {
	Args     *Args
	Retain   bool
	Rebuild  *Trigger
	Restart  *Trigger
	ctx      *context.Context
	cancelFn *context.CancelFunc
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

	return &Run{args, retain, rebuild, restart, nil, nil}
}

// Start spins up the process
func (r *Run) Start(p *Project) {
	if r.Retain {
		ctx, cancelFn := context.WithCancel(context.Background())
		log.Printf("Starting process...")
		startErr := exec.CommandContext(ctx, *p.Target).Start()
		if startErr != nil {
			log.Printf("Failed to start retained process.")
			cancelFn()
			return
		}
		r.ctx = &ctx
		r.cancelFn = &cancelFn
	} else {
		startErr := exec.Command(*p.Target).Start()
		if startErr != nil {
			log.Printf("Failed to start process.")
			return
		}
	}
	log.Printf("Started.")
}

// Stop shuts down the process
func (r *Run) Stop() {
	if r.ctx != nil {
		log.Printf("Stopping process...")
		(*r.cancelFn)()
		r.ctx = nil
		r.cancelFn = nil
		log.Printf("Stopped.")
	}
}
