package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Called if 'info' subcommand is invoked.
func onInfo(c *cli.Context) error {
	fmt.Println("Retrieving information about the font...")

	return nil
}

// Constructs the 'info' subcommand.
func RegisterInfo() *cli.Command {
	return &cli.Command{
		Name:   "info ",
		Usage:  "Prints out information about a font given its identifier in the registry.",
		Action: onInfo,
	}

}
