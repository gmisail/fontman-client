package model

import (
	"fmt"
	"os"

	"fontman/client/pkg/errors"

	"github.com/goccy/go-yaml"
)

type RegistryFile struct {
	Name    string          `yaml:"name"`
	License string          `yaml:"license"`
	Creator string          `yaml:"creator"`
	Styles  []RegistryStyle `yaml:"styles"`
}

type RegistryStyle struct {
	Style string `yaml:"style"`
	Url   string `yaml:"url"`
}

// ReadRegistryFile reads & parses the registry file located at `path`.
func ReadRegistryFile(path string) (*RegistryFile, error) {
	contents, fileErr := os.ReadFile(path)

	if fileErr != nil {
		return nil, &errors.FileError{Message: fmt.Sprintf("Failed to open file with name '%s'", path)}
	}

	var registryFile RegistryFile
	if parseErr := yaml.Unmarshal(contents, &registryFile); parseErr != nil {
		return nil, parseErr
	}

	return &registryFile, nil
}
