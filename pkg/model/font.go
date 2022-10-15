package model

type FontFormat string

const (
	OTF FontFormat = "otf"
	TTF            = "ttf"
	TTC            = "ttc"
)

type FontFamily struct {
	Name     string
	Styles   []FontStyle
	Language string
}

type FontStyle struct {
	Name   string
	Path   string
	Format FontFormat
}
