package model

type ProjectFont struct {
	Name string `yaml:"name"`
	Styles []string `yaml:"styles"`
}

type ProjectFile struct {
	Name string `yaml:"name"` 
	Fonts []ProjectFont `yaml:"fonts"`
}
