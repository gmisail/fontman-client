package model

import (
	"github.com/goccy/go-yaml"

	"os"
)

type ConfigFile struct {
	LocalInstallPath  string `yaml:"local_path"`
	GlobalInstallPath string `yaml:"global_path"`
	RegistryAddress   string `yaml:"registry"`
}

func ReadConfigFile(path string) (*ConfigFile, error) {
	contents, fileErr := os.ReadFile(path)

	if fileErr != nil {
		return nil, fileErr
	}

	configFile := ConfigFile{}
	if parseErr := yaml.Unmarshal(contents, &configFile); parseErr != nil {
		return nil, parseErr
	}

	return &configFile, nil
}
