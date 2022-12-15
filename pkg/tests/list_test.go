package tests

import (
	"fontman/client/pkg/service/fontconfig"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSectionParse(t *testing.T) {
	input := "/test/path/font.ttf: Test Font:style=Regular,Italics"
	sections := fontconfig.SplitSections(input)

	if len(sections) != 3 {
		t.Errorf("Got %d sections, expected 3.", len(sections))
	}
}

func TestLineParse(t *testing.T) {
	input := `/Users/gmisail/Library/Fonts/Ubuntu Mono derivative Powerline Bold Italic.ttf: Ubuntu Mono derivative Powerline:style=Bold Italic
		/opt/X11/share/fonts/75dpi/courO24-ISO8859-1.pcf.gz: Courier:style=Oblique
		/opt/X11/share/fonts/75dpi/helvO14-ISO8859-1.pcf.gz: Helvetica:style=Oblique
		/opt/X11/share/fonts/75dpi/helvO24-ISO8859-1.pcf.gz: Helvetica:style=Oblique
		/opt/X11/share/fonts/misc/18x18ko.pcf.gz: Fixed:style=ko`
	families := fontconfig.ParseListLines(input)
	expectedNames := []string{"Ubuntu Mono derivative Powerline", "Courier", "Helvetica", "Helvetica", "Fixed"}

	assert.Equal(t, 5, len(families))

	for i, name := range expectedNames {
		actualName := families[i].Name
		assert.Equal(t, name, actualName)
	}
}
