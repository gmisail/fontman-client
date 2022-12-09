package util

import (
	"errors"
	"fmt"
	"fontman/client/pkg/model"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

func Cache(verbose bool, force bool) error {
	var flags []string

	if verbose {
		flags = append(flags, "-v")
	}

	if force {
		flags = append(flags, "-f")
	}

	cmd := exec.Command("fc-cache", flags...)
	err := cmd.Run()

	// pipe output into parser

	if err != nil {
		return err
	}
	return nil
}

func ListAll() (string, error) {
	cmd := exec.Command("fc-list", ":", "family", "style", "file")
	output, outputErr := cmd.Output()

	if outputErr != nil {
		return "", outputErr
	}

	return string(output), nil
}

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
	configDir, err := GetFontFolder(isGlobal)
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

// first time setup function. generates fontman config file if it doesn't exist.
func SetupFolders(isGlobal bool) error {

	configPathDir, err := CreateConfigPath()
	if err != nil {
		return err
	}

	configDir := filepath.Join(configPathDir, "config.yml")
	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		// if config doesn't exist, generate it.
		err := GenerateConfig(isGlobal, false)
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

func CheckRoot() bool {
	// os.Geteuid() should be sufficient in detecting if the user running it has root permission
	// SUDO_USER is not reliable as attacker could potentially set the environmental variable
	isRoot := os.Geteuid()
	return isRoot == 0
}
