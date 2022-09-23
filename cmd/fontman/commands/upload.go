package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Called if 'upload' subcommand is invoked.
func onUpload(c *cli.Context) error {
	fmt.Println("upload some font...")

	return nil
}

// Constructs the 'upload' subcommand.
func RegisterUpload() *cli.Command {
	return &cli.Command{
		Name:   "upload",
		Usage:  "Upload a font to the font registry",
		Action: onUpload,
	}
}
