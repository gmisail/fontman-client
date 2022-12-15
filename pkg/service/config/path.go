package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// CreateConfigDirectory creates a configuration directory,
// if it does not exist.
func CreateConfigDirectory() (string, error) {
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

// SetupFolders will setup the necessary files & folders for the program
// to operate. Creates a configuration file if one does not exist.
func SetupFolders(isGlobal bool) error {
	configPathDir, err := CreateConfigDirectory()
	if err != nil {
		return err
	}

	configDir := filepath.Join(configPathDir, "config.yml")
	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		// if config doesn't exist, generate it.
		err := Generate(isGlobal, false)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

// GetFontFolder gets the font directory for the current system.
func GetFontFolder(isGlobal bool) (string, error) {
	if !isGlobal {
		homePath, _ := os.UserHomeDir()
		userPaths := GetLocalFontPaths(homePath)

		// TODO: .local/share/fonts/fontman/_family_name_/
		for i := 0; i < len(userPaths); i++ {
			println(userPaths[i])
			if _, err := os.Stat(userPaths[i]); err == nil {
				// TODO: actually make the fontman directory
				installPath := filepath.Join(userPaths[i], "fontman")
				if err := os.Mkdir(installPath, 0755); !os.IsExist(err) {
					return "", err
				}
				// TODO: confirmation prompt here
				return installPath, nil
			}
		}
		// no valid path exists
		// TODO: should alert the user that no valid path exists
		return "", os.ErrNotExist
	} else {
		// if global, install to /usr/local/share/fonts, etc.
		rootPath := "/"
		rootPaths := GetGlobalFontPaths(rootPath)

		for i := 0; i < len(rootPaths); i++ {
			if _, err := os.Stat(rootPaths[i]); err == nil {
				installPath := filepath.Join(rootPaths[i], "fontman")
				if err := os.Mkdir(installPath, 0755); err != nil {
					return "", err
				}
				return installPath, nil
			}
		}
		// no valid path exists
		// TODO: should alert the user that no valid path exists
		return "", os.ErrNotExist
	}
}
