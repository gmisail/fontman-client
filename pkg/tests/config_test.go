package tests

import (
	"fontman/client/pkg/model"
	"fontman/client/pkg/service/config"
	"testing"
)

func TestProjectFile(t *testing.T) {
	projectFile, _ := model.ReadProjectFile("../config/testdata/valid_fontman.yml")

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
	projectFile, _ := model.ReadProjectFile("../config/testdata/invalid_fontman.yml")

	if projectFile != nil {
		t.Error("Expected error to be thrown.")
	}
}

func TestConfigFile(t *testing.T) {
	configFile, _ := model.ReadConfigFile("../config/testdata/valid_config.yml")
	if configFile == nil {
		t.Error("Failed to load / parse config file.")
	}

	if configFile.LocalInstallPath != "/Users/meow/Library/Fonts" {
		t.Error("Expected name to be '/Users/meow/Library/fonts'")
	}

	if configFile.GlobalInstallPath != "/Library/Fonts" {
		t.Error("Expected name to be '/Library/fonts'")
	}
}

func TestInvalidConfigFile(t *testing.T) {
	configFile, _ := model.ReadConfigFile("../config/testdata/invalid_config.yml")

	if configFile != nil {
		t.Error("Expected error to be thrown.")
	}
}

func TestCreateConfigPath(t *testing.T) {
	configPath, err := config.CreateConfigPath()
	t.Log("configPath is ", configPath)
	if err != nil {
		t.Error("Couldn't create fontman config folder: ", err)
	}
}

func TestGenerateGlobalConfig(t *testing.T) {
	err := config.GenerateConfig(true, false)
	if err != nil {
		t.Error("Generate config failed with error: ", err)
	}
}

func TestGenerateLocalConfig(t *testing.T) {
	err := config.GenerateConfig(false, false)
	if err != nil {
		t.Error("Generate config failed with error: ", err)
	}
}

func TestUpdateGlobalConfig(t *testing.T) {
	err := config.GenerateConfig(true, true)
	if err != nil {
		t.Error("Generate config failed with error: ", err)
	}
}

func TestUpdateLocalConfig(t *testing.T) {
	err := config.GenerateConfig(false, true)
	if err != nil {
		t.Error("Generate config failed with error: ", err)
	}
}

func TestReadConfig(t *testing.T) {
	configFile := model.ConfigFile{}
	configFile, err := config.ReadConfig()
	t.Log("Local install path is ", configFile.LocalInstallPath)
	t.Log("Global install path is ", configFile.GlobalInstallPath)
	t.Log("Registry address is ", configFile.RegistryAddress)
	if err != nil {
		t.Error("Read config failed with error: ", err)
	}
}
