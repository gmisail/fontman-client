package config

import (
	"fontman/client/pkg/model"
	"io/ioutil"
	"log"

	"github.com/goccy/go-yaml"
)

func ReadProjectFile(path string) *model.ProjectFile {
	contents, fileErr := ioutil.ReadFile(path)

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
