package config

import (
	"fontman/client/pkg/model"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

func ReadConfigFile(path string) *model.ConfigFile {
	contents, fileErr := os.ReadFile(path)

	if fileErr != nil {
		log.Println(fileErr)
		return nil
	}

	var configFile model.ConfigFile
	if parseErr := yaml.Unmarshal(contents, &configFile); parseErr != nil {
		log.Println(parseErr)
		return nil
	}

	return &configFile
}

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
