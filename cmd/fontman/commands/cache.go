package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Called if 'cache' subcommand is invoked.
func onCache(c *cli.Context, force bool) error {
	fmt.Println("list some fonts...")
	fmt.Println(force)

	return nil
}

// Constructs the 'cache' subcommand.
func RegisterCache() *cli.Command {
	var force bool

	return &cli.Command{
		Name:  "cache",
		Usage: "Regenerate the cache for currently installed fonts.",
		Action: func(c *cli.Context) error {
			return onCache(c, force)
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "force",
				Aliases:     []string{"f"},
				Usage:       "force",
				Destination: &force,
			},
		},
	}
}
