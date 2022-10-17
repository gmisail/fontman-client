package commands

import (
	"fmt"
	"fontman/client/pkg/util"
	"log"

	"github.com/urfave/cli/v2"
)

// Called if 'uninstall' subcommand is invoked.
func onUninstall(c *cli.Context, style string, ex_style string, global bool) error {
	err := util.SetupFolders(global)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uninstalling font(s)...")

	return nil
}

// Constructs the 'uninstall' subcommand.
func RegisterUninstall() *cli.Command {
	var style string
	var ex_style string
	var global bool

	return &cli.Command{
		Name:  "uninstall",
		Usage: "Uninstall a font given its identifier in the registry.",
		Action: func(c *cli.Context) error {
			return onUninstall(c, style, ex_style, global)
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
				Destination: &ex_style,
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
