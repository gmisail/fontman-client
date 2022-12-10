package errors

import (
	"fmt"

	"github.com/fatih/color"
)

type FileError struct {
	Message string
}

func (e FileError) Error() string {
	return fmt.Sprintf("%s %s", color.RedString("File error:"), e.Message)
}
