package commands

import (
	"fmt"
	"fontman/client/pkg/util"
	"log"

	"github.com/urfave/cli/v2"
)

// Called if 'list' subcommand is invoked.
func onList(c *cli.Context, style string, global bool) error {
	err := util.SetupFolders(global)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listing installed font(s)...")

	return nil
}

// path, style, name, font format (truetype/open),

// Constructs the 'list' subcommand.
func RegisterList() *cli.Command {
	var style string
	var global bool

	return &cli.Command{
		Name:  "list",
		Usage: "List all the installed fonts on the system.",
		Action: func(c *cli.Context) error {
			return onList(c, style, global)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "style",
				Aliases:     []string{"s"},
				Usage:       "List all the fonts matching the specified style.",
				Destination: &style,
			},
			&cli.StringFlag{
				Name:    "exclude",
				Aliases: []string{"e"},
				Usage:   "List all the fonts except for the specified style.",
				// TODO: Should excluded style be stored into a different var?
				Destination: &style,
			},
			&cli.BoolFlag{
				Name:        "global",
				Aliases:     []string{"g"},
				Usage:       "List all the fonts installed in system-wide locations.",
				Destination: &global,
			},
		},
	}
}
