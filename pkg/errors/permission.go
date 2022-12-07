package errors

import (
	"fmt"

	"github.com/fatih/color"
)

type PermissionError struct {
	Message string
}

func (i PermissionError) Error() string {
	return color.RedString(fmt.Sprintf("Permission error: %s", i.Message))
}
