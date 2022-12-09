package service

import (
	"fontman/client/pkg/model"
	"os"

	"github.com/goccy/go-yaml"
)

func ReadRegistryFile(path string) (*model.RegistryFile, error) {
	contents, fileErr := os.ReadFile(path)

	if fileErr != nil {
		return nil, fileErr
	}

	var registryFile model.RegistryFile

	if parseErr := yaml.Unmarshal(contents, &registryFile); parseErr != nil {
		return nil, parseErr
	}

	return &registryFile, nil
}
