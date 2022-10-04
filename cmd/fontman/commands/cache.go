package commands

import (
	"fmt"
	"fontman/client/pkg/util"
	"github.com/urfave/cli/v2"
)

// Called if 'cache' subcommand is invoked.
func onCache(c *cli.Context, verbose bool, force bool) error {
	fmt.Println("cache some font(s)...")
	fmt.Println(verbose, force)
	err := util.Cache(verbose, force)
	if err != nil {
		return err
	}

	return nil
}

// Constructs the 'cache' subcommand.
func RegisterCache() *cli.Command {
	var verbose bool
	var force bool

	return &cli.Command{
		Name:                   "cache",
		Usage:                  "Regenerate the cache for currently installed fonts",
		UseShortOptionHandling: true,
		Action: func(c *cli.Context) error {
			return onCache(c, verbose, force)
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "force",
				Aliases:     []string{"f"},
				Usage:       "Regenerate cache for every font, not just un-cached ones.",
				Destination: &force,
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Usage:       "Regenerate cache for every font, with verbose output.",
				Destination: &verbose,
			},
		},
	}
}
