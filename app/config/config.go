package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
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

func LoadConfig(name string) (Config, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
