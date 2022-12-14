package font

import (
	"fmt"
	"fontman/client/pkg/api"
	"fontman/client/pkg/errors"
	"fontman/client/pkg/service/config"
	"fontman/client/pkg/service/fontconfig"
	"os"
	"path/filepath"
	"strings"
)

func validateFormat(file string) bool {
	fileType := filepath.Ext(file)
	return fileType == ".ttf" || fileType == ".otf" || fileType == ".ttc"
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
	configFile, err := config.Read()
	if err != nil {
		err = config.Generate(isGlobal, false)
		if err != nil {
			return err
		}
	}

	if isGlobal {
		if len(configFile.GlobalInstallPath) == 0 {
			return &errors.InstallationError{
				Message: "Global install path in config is empty.",
			}
		} else {
			installPath = configFile.GlobalInstallPath
		}
	} else {
		if len(configFile.LocalInstallPath) == 0 {
			return &errors.InstallationError{
				Message: "Local install path in config is empty.",
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
	cacheErr := fontconfig.RunCache(false, false)
	if cacheErr != nil {
		return cacheErr
	}

	fmt.Printf("Successfully installed '%s' to '%s'! 🎉\n", fileName, installPath)

	return nil
}

func InstallFromRemote(id string, baseUrl string, isGlobal bool) error {
	font, err := api.GetFontDetails(id, baseUrl)

	if err != nil {
		return err
	}

	// download each style to a file with name: <family>-<style>.<format>
	for _, style := range font.Styles {
		// TODO: replace this with temp path in ~/.fontman

		// get the extension via the URL, i.e. https://fontprovider.com/arial.ttf
		ext := filepath.Ext(style.Url)
		dest := fmt.Sprintf("%s-%s%s", font.Name, style.Type, ext)

		// replace spaces with '-' to prevent any filepath issues
		dest = strings.ReplaceAll(dest, " ", "-")

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
