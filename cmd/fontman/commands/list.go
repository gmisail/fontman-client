package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Called if 'list' subcommand is invoked.
func onList(c *cli.Context) error {
	fmt.Println("list some fonts...")

	return nil
}

// Constructs the 'list' subcommand.
func RegisterList() *cli.Command {
	return &cli.Command {
		Name:  "list",
		Usage: "List all installed fonts.",
		Action: onList,
	}
}
