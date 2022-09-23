package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Called if 'cache' subcommand is invoked.
func onCache(c *cli.Context) error {
	fmt.Println("list some fonts...")

	return nil
}

// Constructs the 'cache' subcommand.
func RegisterCache() *cli.Command {
	return &cli.Command{
		Name:   "cache",
		Usage:  "Regenerate the cache for currently installed fonts.",
		Action: onCache,
	}
}
