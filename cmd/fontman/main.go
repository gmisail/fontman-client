package main

import (
	"fontman/registry/cmd/fontman/commands"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "fontman",
		Usage: "Utility for downloading and managing fonts.",
		Commands: []*cli.Command { 
			commands.RegisterList(),
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
