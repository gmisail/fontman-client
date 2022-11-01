package commands

import (
	"fmt"
	"fontman/client/pkg/model"
	"fontman/client/pkg/parser"
	"fontman/client/pkg/util"
	"log"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

/*
	Given a list of fonts, combine them by family name. Returns a list of the 
	formatted font names as well as a mapping of font name to family.
*/
func showUnique(fonts []*model.FontFamily) ([]string, map[string]model.FontFamily) {
	unique := make(map[string]model.FontFamily)
	names := []string{}

	for _, font := range fonts {
		if family, ok := unique[font.Name]; ok {
			family.Styles = append(family.Styles, font.Styles...)
		} else {
			unique[font.Name] = *font
			names = append(names, strings.Title(font.Name))
		}
	}

	sort.Strings(names)

	return names, unique
}

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

	// get all fonts and combine them based on family
	names, fonts := showUnique(parser.ParseListLines(listOut))

	b := strings.Builder{}
	for _, name := range names {
		font := fonts[name]

		// font name
		b.WriteString(color.YellowString(name))
		
		// if styles exist, print them
		if len(font.Styles) > 0 {
			b.WriteString(" (")

			for i, style := range font.Styles {
				b.WriteString(style.Name)

				if len(font.Styles)-1 != i {
					b.WriteString(", ")
				}
			}

			b.WriteString(")")
		}
		
		// print the line and reset the buffer
		fmt.Println(b.String())
		b.Reset()
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
