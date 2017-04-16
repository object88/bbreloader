package config

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Root     string
	Watch    string
	Target   string
	Build    *Build
	Triggers []*Trigger
}

// SetupConfig reads the configuration and transforms it into living objects
func SetupConfig() (*[]*Config, bool) {
	viper.SetConfigName(".reloader")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("No configuration file found:\n%s\n", err.Error())
		return nil, false
	}

	// configs := viper.Get("configs").([]interface{})[0]
	// log.Printf("%#v\n", configs)

	config := reloaderConfig{}
	viper.Unmarshal(&config)

	log.Printf("Loaded config:\n%#v\n", config)

	c := make([]*Config, len(config.configs))
	for k, v := range config.configs {
		c[k] = parseConfig(v)
	}

	return &c, true
}

func parseConfig(config *config) *Config {
	root, err := filepath.Abs(config.root)
	if err != nil {
		log.Fatalf("Failed to calculate the absolute path: %s\n", err)
		return nil
	}

	watch := root
	if !strings.HasSuffix(watch, string(filepath.Separator)) {
		watch += string(filepath.Separator)
	}
	watch += "..."

	build := parseBuildConfig(config, config.build.steps)

	triggers := make([]*Trigger, len(*config.triggers))
	for i := 0; i < len(triggers); i++ {
		triggers[i] = parseTriggerConfig(config, (*config.triggers)[i])
	}

	return &Config{
		Root:     root,
		Watch:    watch,
		Target:   config.target,
		Build:    build,
		Triggers: triggers,
	}
}
