package main

import (
	"fontman/registry/cmd/fontman/commands"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "fontman",
		Usage: "Utility for downloading and managing fonts.",
		Commands: []*cli.Command{
			commands.RegisterInstall(),
			commands.RegisterUninstall(),
			commands.RegisterInfo(),
			commands.RegisterList(),
			commands.RegisterAuthentication(),
			commands.RegisterSearch(),
			commands.RegisterCache(),
			commands.RegisterUpload(),
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
