package config

import (
	"fontman/client/pkg/model"
	"os"

	"github.com/goccy/go-yaml"
)

func ReadConfigFile(path string) (*model.ConfigFile, error) {
	contents, fileErr := os.ReadFile(path)

	if fileErr != nil {
		return nil, fileErr
	}

	configFile := model.ConfigFile{}
	if parseErr := yaml.Unmarshal(contents, &configFile); parseErr != nil {
		return nil, parseErr
	}

	return &configFile, nil
}

func ReadProjectFile(path string) (*model.ProjectFile, error) {
	contents, fileErr := os.ReadFile(path)

	if fileErr != nil {
		return nil, fileErr
	}

	var projectFile model.ProjectFile

	if parseErr := yaml.Unmarshal(contents, &projectFile); parseErr != nil {
		return nil, parseErr
	}

	return &projectFile, nil
}
