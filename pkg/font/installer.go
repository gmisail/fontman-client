package font

import (
	"fmt"
	"fontman/client/pkg/api"
	"fontman/client/pkg/errors"
	"fontman/client/pkg/util"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func validateFormat(file string) bool {
	fileType := filepath.Ext(file)
	return fileType == ".ttf" || fileType == ".otf" || fileType == ".ttc"
}

// DownloadFrom: downloads file from 'url' and saves it as 'dest`
func DownloadFrom(url string, dest string) error {
	response, resErr := http.Get(url)

	if resErr != nil {
		return resErr
	}

	defer response.Body.Close()

	contents, readErr := io.ReadAll(response.Body)

	if readErr != nil {
		return readErr
	}

	return os.WriteFile(dest, contents, 0777)
}

// InstallFont install a font either locally or globally & regenerate the cache.
func InstallFont(file string, isGlobal bool) error {
	if !validateFormat(file) {
		return &errors.InstallationError{
			Message: fmt.Sprintf("Cannot install font with unsupported format '%s'.", filepath.Ext(file)),
		}
	}

	if _, statErr := os.Stat(file); statErr != nil {
		return &errors.InstallationError{
			Message: fmt.Sprintf("File '%s' does not exist.", file),
		}
	}

	// determine where to install the font
	installPath := ""
	configFile, err := util.ReadConfig()
	if err != nil {
		return err
	}
	if isGlobal {
		if len(configFile.GlobalInstallPath) == 0 {
			return &errors.InstallationError{
				Message: fmt.Sprintf("Global install path in config is empty."),
			}
		} else {
			installPath = configFile.GlobalInstallPath
		}
	} else {
		if len(configFile.LocalInstallPath) == 0 {
			return &errors.InstallationError{
				Message: fmt.Sprintf("Local install path in config is empty."),
			}
		} else {
			installPath = configFile.LocalInstallPath
		}
	}

	fileName := filepath.Base(file)

	// rename the font to the target directory, equivalent to `mv font.ttf ~/dest/font.ttf`
	installErr := os.Rename(file, filepath.Join(installPath, fileName))

	if installErr != nil {
		return installErr
	}

	// after installation, attempt to regenerate the cache
	cacheErr := util.Cache(false, false)
	if cacheErr != nil {
		return cacheErr
	}

	fmt.Printf("Successfully installed '%s' to '%s'!\n", fileName, installPath)

	return nil
}

func InstallFromRemote(id string, isGlobal bool) error {
	configFile, err := util.ReadConfig()
	if err != nil {
		return err
	}
	if len(configFile.RegistryAddress) == 0 {
		return &errors.InstallationError{
			Message: fmt.Sprintf("Registry address is not initialized in config."),
		}
	}

	font, err := api.GetFontDetails(id, configFile.RegistryAddress)
	if err != nil {
		return err
	}

	// download each style to a file with name: <family>-<style>.<format>
	for _, style := range font.Styles {
		// replace spaces with '-' to prevent any filepath issues
		normalizedName := strings.ReplaceAll(font.Name, " ", "-")

		// TODO: replace this with temp path in ~/.fontman
		dest := fmt.Sprintf("%s-%s.%s", normalizedName, style.Type, "ttf")

		// download the font to the working directory
		if err := DownloadFrom(style.Url, dest); err != nil {
			return err
		}

		// install the font to the system
		if err := InstallFont(dest, false); err != nil {
			return err
		}
	}

	return nil
}
