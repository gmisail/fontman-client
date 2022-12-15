package commands

import (
	"fmt"
	"fontman/client/pkg/model"
	"fontman/client/pkg/service/config"
	"fontman/client/pkg/service/registry"

	"github.com/urfave/cli/v2"
)

// Called if 'upload' subcommand is invoked.
func onUpload(c *cli.Context) error {
	if c.Args().Len() == 0 {
		cli.ShowCommandHelp(c, "upload")

		return nil
	}

	fileName := c.Args().Get(0)
	registryFile, readErr := model.ReadRegistryFile(fileName)

	if readErr != nil {
		return readErr
	}

	configFile, configReadErr := config.Read()
	if configReadErr != nil {
		return configReadErr
	}

	if err := registry.UploadRegistryFile(*registryFile, configFile.RegistryAddress); err != nil {
		return err
	}

	fmt.Printf("Successfully uploaded '%s' to the registry.\n", registryFile.Name)

	return nil
}

// Constructs the 'upload' subcommand.
func RegisterUpload() *cli.Command {
	return &cli.Command{
		Name:      "upload",
		Usage:     "Upload a font to the font registry",
		ArgsUsage: "<filename>",
		Action:    onUpload,
	}
}
