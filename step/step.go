package step

import (
	"context"
	"log"
	"path/filepath"
	"strings"

	"github.com/object88/bbreloader/config"
	"github.com/object88/bbreloader/glob"
)

type Config struct {
	Root     string
	Watch    string
	Target   string
	Patterns []*Pattern
}

type Pattern struct {
	Matcher *glob.Matcher
	Steps   []Step
}

// Step describes an increment in the reloader process
type Step interface {
	Run(ctx context.Context) (int, error)
}

func ParseConfig(config *config.Config) *Config {
	patterns := make([]*Pattern, len(config.Patterns))
	for i := 0; i < len(patterns); i++ {
		patterns[i] = parsePatternConfig(config, config.Patterns[i])
	}

	root, err := filepath.Abs(config.Root)
	if err != nil {
		log.Fatalf("Failed to calculate the absolute path: %s\n", err)
		return nil
	}

	watch := root
	if !strings.HasSuffix(watch, string(filepath.Separator)) {
		watch += string(filepath.Separator)
	}
	watch += "..."

	return &Config{
		root,
		watch,
		config.Target,
		patterns,
	}
}

func parsePatternConfig(config *config.Config, pattern *config.PatternConfig) *Pattern {
	m := glob.PreprocessGlobSpec(pattern.Glob)
	steps := make([]Step, len(pattern.Steps))
	for i := 0; i < len(steps); i++ {
		steps[i] = stepConfigToStep(config, pattern.Steps[i])
	}
	return &Pattern{m, steps}
}

func stepConfigToStep(config *config.Config, sc *config.StepConfig) Step {
	switch sc.Type {
	case "build":
		return newStepBuild(config, sc)
	default:
		return newStepCustom()
	}
}
