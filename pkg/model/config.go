package model

type ConfigFile struct {
	LocalInstallPath  string `yaml:"local_path"`
	GlobalInstallPath string `yaml:"global_path"`
	RegistryAddress   string `yaml:"registry"`
}
