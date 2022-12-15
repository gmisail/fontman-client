package commands

import (
	"fmt"
	"fontman/client/pkg/service/config"

	"github.com/urfave/cli/v2"
)

// Called if 'uninstall' subcommand is invoked.
func onUninstall(c *cli.Context, style string, ex_style string, global bool) error {
	err := config.SetupFolders(global)

	if err != nil {
		return err
	}

	fmt.Println("Uninstalling font(s)...")

	return nil
}

// Constructs the 'uninstall' subcommand.
func RegisterUninstall() *cli.Command {
	var style string
	var excludeStyle string
	var global bool

	return &cli.Command{
		Name:  "uninstall",
		Usage: "Uninstall a font given its identifier in the registry.",
		Action: func(c *cli.Context) error {
			return onUninstall(c, style, excludeStyle, global)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "style",
				Aliases:     []string{"s"},
				Usage:       "Specify a style of font to uninstall.",
				Destination: &style,
			},
			&cli.StringFlag{
				Name:        "exclude",
				Aliases:     []string{"e"},
				Usage:       "Uninstall all styles except for the specified.",
				Destination: &excludeStyle,
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
