package config

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Project describes the build and run actions to take for a single project
type Project struct {
	Root     string
	Watch    string
	Target   string
	Build    *Build
	Triggers []*Trigger
}

// SetupConfig reads the configuration and transforms it into living objects
func SetupConfig() (*[]*Project, bool) {
	config := ReloaderMapstructure{}
	viper.Unmarshal(&config)

	log.Printf("Loaded config:\n%#v\n", config)

	c := make([]*Project, len(config.Projects))
	for k, v := range config.Projects {
		c[k] = parseConfig(v)
	}

	return &c, true
}

func parseConfig(project *ProjectMapstructure) *Project {
	root, err := filepath.Abs(project.Root)
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

	triggers := make([]*Trigger, len(*project.Triggers))
	for i := 0; i < len(triggers); i++ {
		triggers[i] = parseTriggerConfig((*project.Triggers)[i])
	}

	return &Project{
		Root:     root,
		Watch:    watch,
		Target:   project.Target,
		Build:    build,
		Triggers: triggers,
	}
}
