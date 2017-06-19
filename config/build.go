package config

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Build contains a series of individual steps necessary to build a project
type Build struct {
	Args      *Args
	PreBuild  []*Step
	PostBuild []*Step
}

func parseBuildConfig(project *ProjectMapstructure, build *BuildMapstructure) *Build {
	if build == nil {
		return &Build{&Args{}, []*Step{}, []*Step{}}
	}
	args := parseArgs(build.Args)
	pre := makeSteps(project, build.PreBuildSteps)
	post := makeSteps(project, build.PostBuildSteps)
	return &Build{args, pre, post}
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
func (s *Build) Run(ctx context.Context, p *Project) (int, error) {

	var tempFileName string
	copy := false
	tempFile, err := ioutil.TempFile("", "tmp")
	if err != nil {
		tempFileName = *p.Target
	} else {
		copy = true
		tempFileName = tempFile.Name()
	}

	completeArgs := s.Args.prependArgs("build", "-o", "tempFileName")

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
