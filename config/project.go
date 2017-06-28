package config

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Project describes the build and run actions to take for a single project
type Project struct {
	Root    string
	Watch   string
	Target  *string
	Builder *Builder
	// Test *Test
	Runner *Runner
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

	runner := parseRun(project, project.Run)

	return &Project{
		Root:    root,
		Watch:   watch,
		Target:  project.Target,
		Builder: build,
		Runner:  runner,
	}
}
