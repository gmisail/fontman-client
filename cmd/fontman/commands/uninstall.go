package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Called if 'uninstall' subcommand is invoked.
func onUninstall(c *cli.Context) error {
	fmt.Println("Uninstalling font(s)...")

	return nil
}

// Constructs the 'uninstall' subcommand.
func RegisterUninstall() *cli.Command {
	var style string
	var global bool

	return &cli.Command{
		Name:   "install",
		Usage:  "Uninstall a font given its identifier in the registry.",
		Action: onUninstall,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "style",
				Aliases:     []string{"s"},
				Usage:       "Specify a style of font to uninstall.",
				Destination: &style,
			},
			&cli.StringFlag{
				Name:    "exclude",
				Aliases: []string{"e"},
				Usage:   "Uninstall all styles except for the specified.",
				// TODO: Should excluded style be stored into a different var?
				Destination: &style,
			},
			&cli.BoolFlag{
				Name:        "global",
				Aliases:     []string{"g"},
				Usage:       "Uninstall fonts in the system-wide locations.",
				Destination: &global,
			},
		},
	}
}
