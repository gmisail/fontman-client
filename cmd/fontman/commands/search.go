package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Called if 'search' subcommand is invoked.
func onSearch(c *cli.Context) error {
	fmt.Println("search for some font...")

	return nil
}

// Constructs the 'search' subcommand.
func RegisterSearch() *cli.Command {
	return &cli.Command{
		Name:   "search",
		Usage:  "Search for a font in the font registry",
		Action: onSearch,
	}
}
