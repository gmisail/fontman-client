package errors

import (
	"fmt"

	"github.com/fatih/color"
)

type InstallationError struct {
	Message string
}

func (i InstallationError) Error() string {
	return fmt.Sprintf("%s %s", color.RedString("Installation error:"), i.Message)
}
