package errors

import (
	"fmt"

	"github.com/fatih/color"
)

type PermissionError struct {
	Message string
}

func (e PermissionError) Error() string {
	return fmt.Sprintf("%s %s", color.RedString("Permission error:"), e.Message)
}
