//go:build darwin
// +build darwin

package config

import (
	"path/filepath"
)

func GetLocalFontPaths(basePath string) []string {
	return []string{
		filepath.Join(basePath, "Library", "Fonts"),
	}
}

func GetGlobalFontPaths(basePath string) []string {
	return GetLocalFontPaths(basePath)
}
