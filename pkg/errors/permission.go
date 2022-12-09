package errors

import (
	"fmt"

	"github.com/fatih/color"
)

type PermissionError struct {
	Message string
}

func (i PermissionError) Error() string {
	return fmt.Sprintf("%s %s", color.RedString("Permission error:"), i.Message)
}
