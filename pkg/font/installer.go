package font

import (
	"fmt"
	"path/filepath"
)

type InstallationError struct {
	message string
}

// TODO: replace with a standard error format
func (i InstallationError) Error() string {
	panic(i.message)
}

func validateFormat(file string) bool {
	fileType := filepath.Ext(file)
	return fileType == "ttf" || fileType == "otf" || fileType == "ttc"
}

func InstallFont(file string) error {
	if !validateFormat(file) {
		return &InstallationError{
			message: fmt.Sprintf("Cannot install font with unsupported format '%s'.", filepath.Ext(file)),
		}
	}
}
