package config

import (
	"context"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Project describes the build and run actions to take for a single project
type Project struct {
	Root   string
	Watch  string
	Target *string
	Build  *Build
	// Test *Test
	Run      *Run
	Context  context.Context
	CancelFn *context.CancelFunc
}

// SetupProjects reads the configuration and transforms it into living objects
func SetupProjects() (*[]*Project, bool) {
	config := ReloaderMapstructure{}
	viper.Unmarshal(&config)

	log.Printf("Loaded config:\n%s\n", config)

	c := make([]*Project, len(config.Projects))
	for k, v := range config.Projects {
		c[k] = parseProject(v)
	}

	return &c, true
}

// Start spins up the process
func (c *Project) Start() {
	// if c.Retain {
	ctx, cancelFn := context.WithCancel(context.Background())
	log.Printf("Starting process...")
	startErr := exec.CommandContext(ctx, *c.Target).Start()
	if startErr != nil {
		log.Printf("Failed to start retained process.")
		cancelFn()
		return
	}
	c.Context = ctx
	c.CancelFn = &cancelFn
	// } else {
	// 	startErr := exec.Command(*c.Target).Start()
	// 	if startErr != nil {
	// 		log.Printf("Failed to start process.")
	// 		return
	// 	}
	// }
	log.Printf("Started.")
}

// Stop shuts down the process
func (c *Project) Stop() {
	if c.Context != nil {
		log.Printf("Stopping process...")
		(*c.CancelFn)()
		c.Context = nil
		c.CancelFn = nil
		log.Printf("Stopped.")
	}
}

func parseProject(project *ProjectMapstructure) *Project {
	root := "."
	if project.Root != nil {
		root = *project.Root
	}
	root, err := filepath.Abs(root)
	if err != nil {
		log.Fatalf("Failed to calculate the absolute path: %s\n", err)
		return nil
	}

	watch := root
	if !strings.HasSuffix(watch, string(filepath.Separator)) {
		watch += string(filepath.Separator)
	}
	watch += "..."

	build := parseBuildConfig(project, project.Build)

	run := parseRun(project, project.Run)

	return &Project{
		Root:   root,
		Watch:  watch,
		Target: project.Target,
		Build:  build,
		Run:    run,
	}
}
