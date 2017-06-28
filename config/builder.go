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

// Builder contains a series of individual steps necessary to build a project
type Builder struct {
	Args          *Args
	PreBuild      []*Step
	PostBuild     []*Step
	tempDir       string
	tempFileIndex int
	restarter     *sync.Restarter
}

// InitializeBuildDirectory will get a temp directory for all the
// build operations.
func (b *Builder) InitializeBuildDirectory() error {
	tempDir, err := ioutil.TempDir("", "")
	if nil != err {
		return err
	}

	b.tempDir = tempDir

	fmt.Printf("Will use build dir '%s'\n", tempDir)
	return nil
}

// DestroyBuildDirectory removes the temp directory.
func (b *Builder) DestroyBuildDirectory() {
	os.Remove(b.tempDir)
}

func parseBuildConfig(project *ProjectMapstructure, build *BuildMapstructure) *Builder {
	r := sync.NewRestarter()
	if build == nil {
		return &Builder{&Args{}, []*Step{}, []*Step{}, "", 0, r}
	}
	args := parseArgs(build.Args)
	pre := makeSteps(project, build.PreBuildSteps)
	post := makeSteps(project, build.PostBuildSteps)
	return &Builder{args, pre, post, "", 0, r}
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
func (b *Builder) Run(p *Project) error {
	b.restarter.Invoke(func(ctx context.Context) {
		b.work(ctx, p)
	})

	return nil
}

func (b *Builder) work(ctx context.Context, p *Project) error {
	tempFileName := fmt.Sprintf("%s/%d.tmp", b.tempDir, b.tempFileIndex)
	b.tempFileIndex++

	// Run pre-build steps
	preStepErr := runSteps(ctx, p, &b.PreBuild)
	if preStepErr != nil {
		return preStepErr
	}

	// Spawn the go build command
	select {
	case <-ctx.Done():
		return context.Canceled
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
		return context.Canceled
	default:
		p.Runner.Stop()
	}

	if p.Target != nil {
		// Move the built file.
		select {
		case <-ctx.Done():
			return context.Canceled
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
		return context.Canceled
	default:
		postStepErr := runSteps(ctx, p, &b.PostBuild)
		if postStepErr != nil {
			return postStepErr
		}
	}

	return nil
}
