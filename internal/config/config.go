package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Environment_Info map[string]string
	Headers          map[string]InputField
	Services         map[string]Service
}

type InputField struct {
	Description string
	Type        string
	Options     []string
	Default     string
	Locked      bool
}

type Service struct {
	Description  string
	Headers      map[string]InputField
	Environments map[string]Environment
}

type Environment struct {
	Headers map[string]InputField
	Url     string
	Paths   map[string]Path
}

type Path struct {
	Description string
	Headers     map[string]InputField
	Variables   map[string]InputField
	Methods     map[string]Method
}

type Method struct {
	Description string
	Headers     map[string]InputField
	Payload     InputField
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
	slog.Debug("configuration loaded successfully", "config", Config)
	return nil
}
