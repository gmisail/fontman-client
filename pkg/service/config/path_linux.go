//go:build linux
// +build linux

package config

import (
	"path/filepath"
)

func GetLocalFontPaths(basePath string) []string {
	return []string{
		filepath.Join(basePath, ".local", "share", "fonts"),
		filepath.Join(basePath, ".fonts"),
	}
}

func GetGlobalFontPaths(basePath string) []string {
	return []string{
		filepath.Join(basePath, "usr", "local", "share", "fonts"),
		filepath.Join(basePath, "usr", "share", "fonts"),
	}
}
