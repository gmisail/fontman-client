package model

const (
	OTF  string = "otf"
	TTF         = "ttf"
	TTC         = "ttc"
	NONE        = "none"
)

type FontFamily struct {
	Name     string
	Styles   []FontStyle
	Language string
}

type FontStyle struct {
	Name   string
	Path   string
	Format string
}
