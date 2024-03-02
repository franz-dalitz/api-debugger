package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Headers  map[string]InputField
	Services map[string]Service
}

type InputField struct {
	Type        string
	Description string
	Locked      bool
	Options     []string
	Default     string
}

type Service struct {
	Headers      map[string]InputField
	Environments map[string]Environment
}

type Environment struct {
	Headers map[string]InputField
	Extends string
	Url     string
	Paths   map[string]Path
}

type Path struct {
	Description string
	Headers     map[string]InputField
	Variables   map[string]InputField
	Methods     map[string]InputField
}

var Config Configuration

func LoadConfig() error {
	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}
	return nil
}
