package commands

import (
	"errors"
	"fontman/client/pkg/model"
	"fontman/client/pkg/service"
	"fontman/client/pkg/util"

	"github.com/urfave/cli/v2"
)

// Called if 'upload' subcommand is invoked.
func onUpload(c *cli.Context) error {
	if c.Args().Len() == 0 {
		return errors.New("File name please")
	}

	fileName := c.Args().Get(0)

	registryFile, readErr := model.ReadRegistryFile(fileName)

	if readErr != nil {
		return readErr
	}

	configFile, configReadErr := util.ReadConfig()
	if configReadErr != nil {
		return configReadErr
	}

	if err := service.UploadRegistryFile(*registryFile, configFile.RegistryAddress); err != nil {
		return err
	}

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
