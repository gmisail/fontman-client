package config_test

import (
	"fontman/client/pkg/config"
	"testing"
)

func TestProjectFile(t *testing.T) {
	projectFile := config.ReadProjectFile("./testdata/valid_fontman.yml")

	if projectFile == nil {
		t.Error("Failed to load / parse project file.")
	}

	if projectFile.Name != "example" {
		t.Error("Expected name to be 'example'")
	}

	if len(projectFile.Fonts) != 1 {
		t.Errorf("Expected file to contain 1 font definition, got %d.", len(projectFile.Fonts))
	}
}

func TestInvalidProjectFile(t *testing.T) {
	projectFile := config.ReadProjectFile("./testdata/invalid_fontman.yml")

	if projectFile != nil {
		t.Error("Expected error to be thrown.")
	}
}
