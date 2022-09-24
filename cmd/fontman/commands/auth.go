package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Called if 'auth' subcommand is invoked.
func onAuthenticate(c *cli.Context) error {
	fmt.Println("authenticate")

	return nil
}

// Constructs the 'auth' subcommand.
func RegisterAuthentication() *cli.Command {
	return &cli.Command{
		Name:   "auth",
		Usage:  "Authenticate or unauthenticate client",
		Action: onAuthenticate,
	}
}
