package commands

import (
	"fmt"
	"fontman/client/pkg/parser"
	"fontman/client/pkg/util"
	"log"
	"sort"
	"strings"

	"github.com/urfave/cli/v2"
)

// Called if 'list' subcommand is invoked.
func onList(c *cli.Context, style string, global bool) error {
	err := util.SetupFolders(global)

	if err != nil {
		log.Fatal(err)

		return err
	}

	listOut, listOutErr := util.ListAll()

	if listOutErr != nil {
		log.Fatal(listOutErr)

		return listOutErr
	}

	allFonts := parser.ParseListLines(listOut)

	sort.Slice(allFonts, func(i, j int) bool {
		return allFonts[i].Name < allFonts[j].Name
	})

	b := strings.Builder{}
	for _, font := range allFonts {
		b.Reset()

		b.WriteString(font.Name)
		b.WriteString(" ")
		b.WriteString(font.Language)
		b.WriteString(" (")

		for i, style := range font.Styles {
			b.WriteString(style.Name)

			if len(font.Styles)-1 != i {
				b.WriteString(", ")
			}
		}

		b.WriteString(")")

		fmt.Println(b.String())
	}

	return nil
}

// path, style, name, font format (truetype/open),

// Constructs the 'list' subcommand.
func RegisterList() *cli.Command {
	var style string
	var global bool

	return &cli.Command{
		Name:  "list",
		Usage: "List all the installed fonts on the system.",
		Action: func(c *cli.Context) error {
			return onList(c, style, global)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "style",
				Aliases:     []string{"s"},
				Usage:       "List all the fonts matching the specified style.",
				Destination: &style,
			},
			&cli.StringFlag{
				Name:    "exclude",
				Aliases: []string{"e"},
				Usage:   "List all the fonts except for the specified style.",
				// TODO: Should excluded style be stored into a different var?
				Destination: &style,
			},
			&cli.BoolFlag{
				Name:        "global",
				Aliases:     []string{"g"},
				Usage:       "List all the fonts installed in system-wide locations.",
				Destination: &global,
			},
		},
	}
}
