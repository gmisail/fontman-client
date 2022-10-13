package parser

import (
	"fontman/client/pkg/model"
	"strings"
)

func SplitSections(line string) []string {
	sections := strings.Split(line, ":")

	for i, section := range sections {
		sections[i] = strings.TrimSpace(section)
	}

	return sections
}

func ParseListLine(line string) model.FontFamily {
	sections := SplitSections(line)

	// path := sections[0]
	styles := ParseStyles(sections[2])

	return model.FontFamily{
		Name: "family",
		Styles: styles,
		Language: "en",
	}
}

func ParseFontPath(path string) {

}

func ParseStyles(path string, styles string) []model.FontStyle {
	styleNames := strings.Split(strings.Trim(styles, "style="), ",")

	
}
