package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/gmisail/fontman-client/pkg/model"
	"github.com/gmisail/fontman-client/pkg/service/config"
	"github.com/gmisail/fontman-client/pkg/service/fontconfig"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/*
Given a list of fonts, combine them by family name. Returns a list of the
formatted font names as well as a mapping of font name to family.
*/
func showUnique(fonts []*model.FontFamily) ([]string, map[string][]string) {
	// family => list of unique styles
	unique := make(map[string][]string)

	// family => "set" of styles, ensures that styles are unique before appending
	uniqueStyles := make(map[string]map[string]struct{})

	// unique font family names
	names := []string{}

	// convert to title based on English
	caser := cases.Title(language.English)

	// for each font, find all *unique* styles & combine common font families.
	for _, font := range fonts {
		// if the font already exists, don't bother re-registering
		if family, ok := unique[font.Name]; ok {
			for _, style := range font.Styles {
				// only add style if it is unique, i.e. it hasn't been registered yet
				if _, hasStyle := uniqueStyles[font.Name][style.Name]; !hasStyle {
					unique[font.Name] = append(family, style.Name)
					uniqueStyles[font.Name][style.Name] = struct{}{}
				}
			}
		} else {
			// register new font
			unique[font.Name] = []string{}
			uniqueStyles[font.Name] = make(map[string]struct{})

			for _, style := range font.Styles {
				unique[font.Name] = append(unique[font.Name], style.Name)
				uniqueStyles[font.Name][style.Name] = struct{}{}
			}

			names = append(names, caser.String(font.Name))
		}
	}

	// sort each of the styles so that order is deterministic
	for familyName := range unique {
		sort.Strings(unique[familyName])
	}

	// sort all font names
	sort.Strings(names)

	return names, unique
}

// Called if 'list' subcommand is invoked.
func onList(c *cli.Context, style string, global bool) error {
	err := config.SetupFolders(global)

	if err != nil {
		return err
	}

	listOut, listOutErr := fontconfig.ListAll()

	if listOutErr != nil {
		return listOutErr
	}

	// get all fonts and combine them based on family
	names, fonts := showUnique(fontconfig.ParseListLines(listOut))

	b := strings.Builder{}
	for _, name := range names {
		font := fonts[name]

		// font name
		b.WriteString(color.YellowString(name))

		// if styles exist, print them
		if len(font) > 0 {
			b.WriteString(" (")

			for i, style := range font {
				b.WriteString(style)

				if len(font)-1 != i {
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
