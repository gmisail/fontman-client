package util

// we use log for now; should move to a more robust way when commands are finalized.
import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
		log.Fatal(err)
		return err
	}
	return nil
}

func SetupFolders(global bool) error {
	// create fontman folder in .config if it doesn't exist
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
		return err
	}

	configDir += "/.fontman"
	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(configDir, 0755)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	// use global flag to determine if we need to check and create $HOME/.fontman folder
	// if global is true, we assume system font folders such as /usr/share/fonts exists
	if !global {
		fontmanDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
			return err
		}

		fontmanDir += "/.fontman"

		if _, err := os.Stat(fontmanDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(fontmanDir, 0755)
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}

	return nil
}

func GetInstallationPath() string {
	dir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(dir, "/.fontman")
}

func CheckRoot() bool {
	// os.Geteuid() should be sufficient in detecting if the user running it has root permission
	// SUDO_USER is not reliable as attacker could potentially set the environmental variable
	isRoot := os.Geteuid()
	return isRoot == 0
}
