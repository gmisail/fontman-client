package parser

import (
	//	"fmt"
	"fontman/client/pkg/model"
	"path/filepath"
	"strings"
)

/*
Split the line into three sections: path, family name, and styles

<path>:<family>,<family-alt>,...:<style>,<style>,...
*/
func SplitSections(line string) []string {
	sections := strings.Split(strings.TrimSpace(line), ":")

	for i, section := range sections {
		sections[i] = strings.TrimSpace(section)
	}

	return sections
}

/*
Given a string of lines (each line -> one font family definition), parse
them and return a list of FontFamily models
*/
func ParseListLines(lines string) []*model.FontFamily {
	var families []*model.FontFamily

	for _, line := range strings.Split(lines, "\n") {
		currentFamily := ParseListLine(line)

		if currentFamily != nil {
			families = append(families, ParseListLine(line))
		}
	}

	return families
}

/*
Split all families into a list of names
*/
func ParseFamilyNames(fullName string) []string {
	return strings.Split(fullName, ",")
}

/*
Parse a line into a FontFamily model
*/
func ParseListLine(line string) *model.FontFamily {
	if len(line) == 0 {
		return nil
	}

	sections := SplitSections(line)

	path := sections[0]
	fontFamilyNames := ParseFamilyNames(sections[1])
	fontFormat := filepath.Ext(path)
	styles := []model.FontStyle{}

	if len(sections) > 2 {
		styles = ParseStyles(path, fontFormat, sections[2])
	}

	// TODO: by default, we're only parsing the English (en) version. Potentially we could load in all languages.
	return &model.FontFamily{
		Name:     fontFamilyNames[0],
		Styles:   styles,
		Language: "en",
	}
}

/*
Parse a style section into FontStyle models
*/
func ParseStyles(filePath string, fileFormat string, styles string) []model.FontStyle {
	styleNames := strings.Split(strings.TrimLeft(styles, "style="), ",")
	var fontStyles []model.FontStyle

	for _, name := range styleNames {
		fontStyles = append(fontStyles, model.FontStyle{Name: name, Path: filePath, Format: fileFormat})
	}

	return fontStyles
}
