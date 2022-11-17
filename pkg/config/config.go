package config

import (
	"fontman/client/pkg/model"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

func ReadProjectFile(path string) *model.ProjectFile {
	contents, fileErr := os.ReadFile(path)

	if fileErr != nil {
		log.Print(fileErr)

		return nil
	}

	var projectFile model.ProjectFile

	if parseErr := yaml.Unmarshal(contents, &projectFile); parseErr != nil {
		log.Print(parseErr)

		return nil
	}

	return &projectFile
}
