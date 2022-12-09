package model

type RegistryFile struct {
	Name    string          `yaml:"name"`
	License string          `yaml:"license"`
	Creator string          `yaml:"creator"`
	Styles  []RegistryStyle `yaml:"styles"`
}

type RegistryStyle struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}
