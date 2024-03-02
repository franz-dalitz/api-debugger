package config

import (
	"flag"
	"os"
)

// var declarations that also initialize default values
var (
	ConfigFile string = "/etc/api-debugger/config.yaml"
	LogLevel   string = "INFO"
)

// only assigns value to string variable if env var is set
func assignPresentEnv(variable *string, key string) {
	value, present := os.LookupEnv(key)
	if present {
		*variable = value
	}
}

func LoadSettings() {
	// read env vars
	assignPresentEnv(&ConfigFile, "ConfigFile")
	assignPresentEnv(&LogLevel, "LogLevel")

	// read cli flags and if present override values set via env vars
	flag.StringVar(&ConfigFile, "config-file", ConfigFile, "path to the config file")
	flag.StringVar(&LogLevel, "log-level", LogLevel, "desired minimum emitted level of logs")

	flag.Parse()

	// apply settings that were just read
	ApplySlogSettings()
}
