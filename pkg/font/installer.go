package font

import (
	"fmt"
	"fontman/client/pkg/api"
	"fontman/client/pkg/errors"
	"fontman/client/pkg/util"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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

	contents, readErr := ioutil.ReadAll(response.Body)

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
	installPath, err := util.GetFontFolder(isGlobal)
	if err != nil {
		return err
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

	fmt.Printf("Successfully installed '%s' to %s! 🎉\n", fileName, installPath)

	return nil
}

func InstallFromRemote(id string) error {
	font, err := api.GetFontDetails(id)

	if err != nil {
		return err
	}

	// download each style to a file with name: <family>-<style>.<format>
	for _, style := range font.Styles {
		dest := fmt.Sprintf("%s-%s.%s", font.Name, style.Type, "ttf")

		if err := DownloadFrom(style.Url, dest); err != nil {
			return err
		}
	}

	return nil
}
