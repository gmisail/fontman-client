package parser_test

import (
	"testing"
	"fontman/client/pkg/parser"
)

func TestLineParse(t *testing.T) {
	input := "/test/path/font.ttf: Test Font:style=Regular,Italics"
	sections := parser.SplitSections(input)

	if len(sections) != 3 {
		t.Errorf("Got %d sections, expected 3.", len(sections))
	}
}
