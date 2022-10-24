package commands

import (
	"fontman/client/pkg/font"
	"github.com/urfave/cli/v2"
)

// Called if 'install' subcommand is invoked.
func onInstall(c *cli.Context, style string, global bool) error {
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
	var style string
	var global bool

	return &cli.Command{
		Name:  "install",
		Usage: "Install a font given its identifier in the registry.",
		Action: func(context *cli.Context) error {
			return onInstall(context, style, global)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "style",
				Aliases:     []string{"s"},
				Usage:       "Specify a style of font to install.",
				Destination: &style,
			},
			&cli.StringFlag{
				Name:    "exclude",
				Aliases: []string{"e"},
				Usage:   "Install all styles except for the specified.",
				// TODO: Should excluded style be stored into a different var?
				Destination: &style,
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
