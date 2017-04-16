package config

type reloaderConfig struct {
	configs []*config `mapstructure:"configs"`
}

type config struct {
	root   string `mapstructure:"root"`
	target string `mapstructure:"target"`
	build  *struct {
		steps []*stepConfig `mapstructure:"steps"`
	} `mapstructure:"patterns"`
	triggers *[]*triggerConfig `mapstructure:"triggers"`
}

type stepConfig struct {
	t       string    `mapstructure:"type"`
	command *string   `mapstructure:"command"`
	args    *[]string `mapstructure:"args"`
}

type triggerConfig struct {
	action string `mapstructure:"action"`
	glob   string `mapstructure:"glob"`
}
