package model

import (
	"github.com/goccy/go-yaml"
	"os"
)

type ProjectFont struct {
	Name   string   `yaml:"name"`
	Styles []string `yaml:"styles"`
}

type ProjectFile struct {
	Name  string        `yaml:"name"`
	Fonts []ProjectFont `yaml:"fonts"`
}

func ReadProjectFile(path string) (*ProjectFile, error) {
	contents, fileErr := os.ReadFile(path)

	if fileErr != nil {
		return nil, fileErr
	}

	var projectFile ProjectFile

	if parseErr := yaml.Unmarshal(contents, &projectFile); parseErr != nil {
		return nil, parseErr
	}

	return &projectFile, nil
}
