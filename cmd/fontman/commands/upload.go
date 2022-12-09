package commands

import (
	"errors"
	"fmt"
	"fontman/client/pkg/service"

	"github.com/urfave/cli/v2"
)

// Called if 'upload' subcommand is invoked.
func onUpload(c *cli.Context) error {
	if c.Args().Len() == 0 {
		return errors.New("File name please")
	}

	fileName := c.Args().Get(0)

	file, _ := service.ReadRegistryFile(fileName)

	fmt.Println(file)

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
