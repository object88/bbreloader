package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/object88/sync"
)

// Build contains a series of individual steps necessary to build a project
type Build struct {
	Args          *Args
	PreBuild      []*Step
	PostBuild     []*Step
	tempDir       string
	tempFileIndex int
	restarter     *sync.Restarter
}

// InitializeBuildDirectory will get a temp directory for all the
// build operations.
func (b *Build) InitializeBuildDirectory() error {
	tempDir, err := ioutil.TempDir("", "")
	if nil != err {
		return err
	}

	b.tempDir = tempDir

	fmt.Printf("Will use build dir '%s'\n", tempDir)
	return nil
}

// DestroyBuildDirectory removes the temp directory.
func (b *Build) DestroyBuildDirectory() {
	os.Remove(b.tempDir)
}

func parseBuildConfig(project *ProjectMapstructure, build *BuildMapstructure) *Build {
	r := sync.NewRestarter()
	if build == nil {
		return &Build{&Args{}, []*Step{}, []*Step{}, "", 0, r}
	}
	args := parseArgs(build.Args)
	pre := makeSteps(project, build.PreBuildSteps)
	post := makeSteps(project, build.PostBuildSteps)
	return &Build{args, pre, post, "", 0, r}
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
	b.restarter.Invoke(func(ctx context.Context) {
		b.work(ctx, p)
	})

	return 0, nil
}

func (b *Build) work(ctx context.Context, p *Project) error {
	tempFileName := fmt.Sprintf("%s/%d.tmp", b.tempDir, b.tempFileIndex)
	b.tempFileIndex++

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
		p.Run.Stop()
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
