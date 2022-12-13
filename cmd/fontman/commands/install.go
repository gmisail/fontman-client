package commands

import (
	"errors"
	"fmt"
	"fontman/client/pkg/service"
	"fontman/client/pkg/service/config"
	"fontman/client/pkg/service/font"
	"strings"

	"fontman/client/pkg/api"
	fontmanErr "fontman/client/pkg/errors"
	"fontman/client/pkg/model"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func selectFromList(options []model.RemoteFontFamily) string {
	// one option? Don't bother showing a list
	if len(options) == 1 {
		return options[0].Id
	}

	view := strings.Builder{}

	for i, option := range options {
		view.WriteString(fmt.Sprintf("%d) %s [%s]\n", i+1, option.Name, option.Id))
	}

	view.WriteString(fmt.Sprintf("\nSelect an option to install [1 - %d]:", len(options)))
	fmt.Println(view.String())

	var selection int
	numScanned, _ := fmt.Scanf("%d", &selection)

	if numScanned == 0 {
		return options[0].Id
	}

	// the values are 1-indexed to look more normal, so we need to adjust for this
	selection -= 1

	if selection < 0 || selection >= len(options) {
		return ""
	}

	return options[selection].Id
}

// installRemote: fetches options, shows a selection view, and then installs based on user selection
func installRemote(fileName string, global bool) error {
	configFile, err := config.ReadConfig()

	if err != nil {
		err = config.GenerateConfig(global, false)
		if err != nil {
			return err
		}

		// re-read the config to make sure RegistryAddress can be checked
		configFile, err = config.ReadConfig()
		if err != nil {
			return err
		}
	}

	if len(configFile.RegistryAddress) == 0 {
		return &fontmanErr.InstallationError{
			Message: "Registry address is not initialized in config.",
		}
	}

	options, optionErr := api.GetFontOptions(fileName, configFile.RegistryAddress)
	var id string

	if optionErr != nil {
		return optionErr
	}

	// more than one option? Ask the user which they want to install
	if len(options) >= 1 {
		id = selectFromList(options)

		if len(id) == 0 {
			return errors.New(fmt.Sprintf("Invalid option selected."))
		}
	} else {
		// no options, throw error
		return errors.New(fmt.Sprintf("No font found with name '%s'", fileName))
	}

	return font.InstallFromRemote(id, configFile.RegistryAddress, global)
}

// Called if 'install' subcommand is invoked.
func onInstall(c *cli.Context, style string, excludeStyle string, global bool) error {
	// if global flag is set, but user doesn't have permission
	if global && !service.CheckRoot() {
		return errors.New("no root permission; run it again with sudo")
	}

	fileName := c.Args().Get(0)

	// no arguments: install from local `fontman.yml` file
	if len(fileName) == 0 {
		// TODO: add multiple options, i.e. fontman.yaml, FontmanFile
		project, projectErr := model.ReadProjectFile("fontman.yml")

		if projectErr != nil {
			return projectErr
		}

		// for each font, try to install from remote
		for _, projectFont := range project.Fonts {
			if err := installRemote(projectFont.Name, global); err != nil {
				fmt.Println(err.Error())
			}
		}

		return nil
	}

	// test.ttf: local file, test: remote file
	ext := filepath.Ext(fileName)

	// if there's an extension, then we're trying to install from local
	if len(ext) != 0 {
		return font.InstallFont(fileName, global)
	}

	return installRemote(fileName, global)
}

// RegisterInstall Constructs the 'install' subcommand.
func RegisterInstall() *cli.Command {
	// TODO: style/ex_style should be arrays of strings; look into how the lib handles multi-parameter argument
	var style string
	var excludeStyle string
	var global bool

	return &cli.Command{
		Name:  "install",
		Usage: "Install a font given its identifier in the registry.",
		Action: func(c *cli.Context) error {
			err := onInstall(c, style, excludeStyle, global)

			if err != nil {
				return err
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "style",
				Aliases:     []string{"s"},
				Usage:       "Specify a style of font to install.",
				Destination: &style,
			},
			&cli.StringFlag{
				Name:        "exclude",
				Aliases:     []string{"e"},
				Usage:       "Install all styles except for the specified.",
				Destination: &excludeStyle,
			},
			&cli.BoolFlag{
				Name:        "global",
				Aliases:     []string{"g"},
				Usage:       "Install fonts to system-wide locations.",
				Destination: &global,
			},
		},
	}
}
