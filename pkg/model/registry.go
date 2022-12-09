package model

import (
	"os"

	"github.com/goccy/go-yaml"
)

type RegistryFile struct {
	Name    string          `yaml:"name"`
	License string          `yaml:"license"`
	Creator string          `yaml:"creator"`
	Styles  []RegistryStyle `yaml:"styles"`
}

type RegistryStyle struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

func ReadRegistryFile(path string) (*RegistryFile, error) {
	contents, fileErr := os.ReadFile(path)

	if fileErr != nil {
		return nil, fileErr
	}

	var registryFile RegistryFile

	if parseErr := yaml.Unmarshal(contents, &registryFile); parseErr != nil {
		return nil, parseErr
	}

	return &registryFile, nil
}
