package commands

import (
	"fmt"
  "fontman/client/pkg/font"
	"fontman/client/pkg/util"
	"log"
	"github.com/urfave/cli/v2"
)

// Called if 'install' subcommand is invoked.
func onInstall(c *cli.Context, style string, excludeStyle string, global bool) error {
	fileName := c.Args().Get(0)

	// no arguments, install from local file
	if len(fileName) == 0 {
		// TODO: return font.InstallFromFile()
		return nil
	}

	// TODO: try to install from remote

	return font.InstallFont(fileName, global)
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
