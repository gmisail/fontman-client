package commands

import (
	"errors"
	"fmt"
	"fontman/client/pkg/font"
	"fontman/client/pkg/util"
	"path/filepath"

	"github.com/urfave/cli/v2"
	// "log"
)

// Called if 'install' subcommand is invoked.
func onInstall(c *cli.Context, style string, excludeStyle string, global bool) error {
	// if global flag is set, but user doesn't have permission
	if global && !util.CheckRoot() {
		return errors.New("no root permission; run it again with sudo")
	}

	fileName := c.Args().Get(0)

	// no arguments: install from local `fontman.yml` file
	if len(fileName) == 0 {
		fmt.Println("fetching from fontman.yml file...")
		return nil
	}

	// test.ttf: local fille, test: remote file
	ext := filepath.Ext(fileName)

	// if there's an extensiuon, then we're trying to install from loccal
	if len(ext) != 0 {
		return font.InstallFont(fileName, global)
	}

	return font.InstallFromRemote("608308d7-5f72-45f4-9776-a99a692703d6")
}

// Constructs the 'install' subcommand.
func RegisterInstall() *cli.Command {
	// TODO: style/ex_style should be arrays of strings; look into how the lib handles multi-parameter argument
	var style string
	var excludeStyle string
	var global bool

	return &cli.Command{
		Name:  "install",
		Usage: "Install a font given its identifier in the registry.",
		Action: func(c *cli.Context) error {
			return onInstall(c, style, excludeStyle, global)
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
