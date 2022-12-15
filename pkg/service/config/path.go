package config

import (
	"errors"
	"os"
	"path/filepath"
)

// first time setup function. generates fontman config file if it doesn't exist.
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

func GetFontFolder(isGlobal bool) (string, error) {
	if !isGlobal {
		homePath, _ := os.UserHomeDir()
		userPaths := []string{
			filepath.Join(homePath, ".local", "share", "fonts"),
			filepath.Join(homePath, ".fonts"),
			// TODO: os detection at compile time - build constraints
			filepath.Join(homePath, "Library", "Fonts"),
		}
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
		rootPaths := []string{
			filepath.Join(rootPath, "usr", "local", "share", "fonts"),
			filepath.Join(rootPath, "usr", "share", "fonts"),
			// TODO: os detection at compile time - build constraints
			filepath.Join(rootPath, "Library", "Fonts"),
		}
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
