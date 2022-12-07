package errors

import (
	"fmt"

	"github.com/fatih/color"
)

type InstallationError struct {
	Message string
}

func (i InstallationError) Error() string {
	return color.RedString(fmt.Sprintf("Installation error: %s", i.Message))
}
