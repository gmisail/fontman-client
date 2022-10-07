package main

import (
	"fontman/client/cmd/fontman/commands"
	"fontman/client/pkg/util"
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

	err := util.SetupFolders()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
