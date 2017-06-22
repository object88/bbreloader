package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
)

var tempDir string
var tempFileIndex int

// InitializeBuildDirectory will get a temp directory for all the
// build operations.
func InitializeBuildDirectory() error {
	var err error
	tempDir, err = ioutil.TempDir("", "")
	if nil != err {
		return err
	}

	fmt.Printf("Will use build dir '%s'\n", tempDir)
	return nil
}

// DestroyBuildDirectory removes the temp directory.
func DestroyBuildDirectory() {
	os.Remove(tempDir)
}

// Build contains a series of individual steps necessary to build a project
type Build struct {
	Args      *Args
	PreBuild  []*Step
	PostBuild []*Step
	cancelFn  *context.CancelFunc
	mutex     *sync.Mutex
}

func parseBuildConfig(project *ProjectMapstructure, build *BuildMapstructure) *Build {
	m := sync.Mutex{}
	if build == nil {
		return &Build{&Args{}, []*Step{}, []*Step{}, nil, &m}
	}
	args := parseArgs(build.Args)
	pre := makeSteps(project, build.PreBuildSteps)
	post := makeSteps(project, build.PostBuildSteps)
	return &Build{args, pre, post, nil, &m}
}

func makeSteps(project *ProjectMapstructure, steps *[]StepMapstructure) []*Step {
	if steps == nil || len(*steps) == 0 {
		return []*Step{}
	}

	count := len(*steps)
	result := make([]*Step, count)
	for i := 0; i < count; i++ {
		result[i] = stepConfigToStep(project, &(*steps)[i])
	}

	return result
}

// Run executes the step with an interruptable context
func (b *Build) Run(p *Project) (int, error) {
	earlyCancel := false

	// Lock the build file and check for a cancellation function
	b.mutex.Lock()
	if *b.cancelFn != nil {
		(*b.cancelFn)()
		b.cancelFn = nil
		earlyCancel = true
	}

	// TODO: fix!
	// This is going to cause some trouble.  If there was already a build running,
	// and we come along and cancel it, then we don't want to use `cancelFn` below
	// (or clear it out from b.mutex), because it will have changed.
	ctx, cancelFn := context.WithCancel(context.Background())
	b.cancelFn = &cancelFn

	b.mutex.Unlock()

	// Do the work
	go b.work(ctx, p)

	// Clean up
	b.mutex.Lock()
	cancelFn()
	if !earlyCancel {
		b.cancelFn = nil
	}
	b.mutex.Unlock()

	return 0, nil
}

func (b *Build) work(ctx context.Context, p *Project) error {
	tempFileName := fmt.Sprintf("%s/%d.tmp", tempDir, tempFileIndex)
	tempFileIndex++

	// Run pre-build steps
	runSteps(ctx, p, &b.PreBuild)

	// Spawn the go build command
	select {
	case <-ctx.Done():
		return nil
	default:
		completeArgs := b.Args.prependArgs("build", "-o", tempFileName)
		cmd := exec.CommandContext(ctx, "go", *completeArgs...)
		cmd.Dir = p.Root
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout

		err := cmd.Run()
		if err != nil {
			log.Fatalf("Build command failed: %s\n", err.Error())
			return err
		}
	}

	// Stop any previously running instance
	select {
	case <-ctx.Done():
		return nil
	default:
		p.Stop()
	}

	if p.Target != nil {
		// Move the built file.
		select {
		case <-ctx.Done():
			return nil
		default:
			// WARNING; there are some issues with this strategy:
			// https://stackoverflow.com/questions/30385225/in-go-is-there-an-os-independent-way-to-atomically-overwrite-a-file
			linkErr := os.Rename(tempFileName, *p.Target)
			if linkErr != nil {
				// Crap.
				fmt.Printf("%s\n", linkErr.Error())
				return linkErr
			}
		}
	}

	// Run post-build steps
	select {
	case <-ctx.Done():
		return nil
	default:
		runSteps(ctx, p, &b.PostBuild)
	}

	return nil
}
