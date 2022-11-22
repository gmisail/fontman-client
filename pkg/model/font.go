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

type RemoteFontFamily struct {
	Id     string            `json:"id"`
	Name   string            `json:"name"`
	Styles []RemoteFontStyle `json:"styles"`
}

type RemoteFontStyle struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}
