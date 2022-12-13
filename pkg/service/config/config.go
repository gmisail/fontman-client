package config

import (
	"errors"
	"fmt"
	"fontman/client/pkg/model"
	"fontman/client/pkg/service/path"
	"github.com/goccy/go-yaml"
	"os"
	"path/filepath"
)

// Create .config/fontman and return it; if it already exists, return it.
func CreateConfigPath() (string, error) {
	// home folder should _really_ exist
	configPath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// we concatenate .config instead of using OS-agnostic UserConfigDir because OSX doesn't allow
	// creating a file of perm 755 in ~/Library/Application Support anymore.
	configPath = filepath.Join(configPath, ".config")
	// check if .config exists
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(configPath, 0755)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		// if stat failed not because folder doesn't exist
		return "", err
	}

	configPath = filepath.Join(configPath, "fontman")
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("fontman config folder doesn't exist; creating...")
		err := os.Mkdir(configPath, 0755)
		if err != nil {
			return "", err
		}
		return configPath, nil
	} else if err != nil {
		// if stat failed and it's not because folder doesn't exist
		return "", err
	} else {
		// if fontman folder already exists
		return configPath, nil
	}
}

func ReadConfig() (model.ConfigFile, error) {
	// ReadConfig assumes that the config file already exists.
	configFile := model.ConfigFile{}
	configPath, err := CreateConfigPath()
	if err != nil {
		return configFile, err
	}
	configPath = filepath.Join(configPath, "config.yml")
	if _, err := os.Stat(configPath); err != nil {
		return configFile, err
	} else {
		// if config exists, read it
		contents, fileErr := os.ReadFile(configPath)
		if fileErr != nil {
			return configFile, fileErr
		}

		// parse the config and return it
		if parseErr := yaml.Unmarshal(contents, &configFile); parseErr != nil {
			return configFile, parseErr
		}
		return configFile, nil
	}
}

// GenerateConfig only initializes one of the two InstallPath fields
func GenerateConfig(isGlobal bool, update bool) error {
	configDir, err := path.GetFontFolder(isGlobal)
	// if no valid folder found, return err
	if err != nil {
		return err
	}

	configFile := model.ConfigFile{}
	if update {
		configFile, err = ReadConfig()
		if err != nil {
			return err
		}
	}

	if isGlobal {
		configFile.GlobalInstallPath = configDir
	} else {
		configFile.LocalInstallPath = configDir
	}

	configFile.RegistryAddress = "https://fontman-registry.up.railway.app"

	configData, err := yaml.Marshal(configFile)
	if err != nil {
		return err
	}

	configFilePath, err := CreateConfigPath()
	if err != nil {
		return err
	}

	configFilePath = filepath.Join(configFilePath, "config.yml")
	err = os.WriteFile(configFilePath, configData, 0755)
	if err != nil {
		return err
	}

	return nil
}
