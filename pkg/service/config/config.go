package config

import (
	"os"
	"path/filepath"

	"github.com/gmisail/fontman-client/pkg/model"

	"github.com/goccy/go-yaml"
)

// Read will attempt to read the user's configuration
// file (`~/.config/fontman/config.yml`), if it exists.
func Read() (model.ConfigFile, error) {
	// ReadConfig assumes that the config file already exists.
	configFile := model.ConfigFile{}
	configPath, err := CreateConfigDirectory()
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

// GenerateConfig will generate a configuration file with
// sane defaults (only initializes one of the two InstallPath fields).
func Generate(isGlobal bool, update bool) error {
	configDir, err := GetFontFolder(isGlobal)
	// if no valid folder found, return err
	if err != nil {
		return err
	}

	configFile := model.ConfigFile{}
	if update {
		configFile, err = Read()
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

	configFilePath, err := CreateConfigDirectory()
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
