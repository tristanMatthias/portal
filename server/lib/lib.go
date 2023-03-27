package lib

import (
	"os"
	"path/filepath"
)

func ConfigPath(fp ...string) string {
	// Join all paths with the config path
	return filepath.Join(
		append([]string{
			os.Getenv("HOME"),
			".portal",
		}, fp...)...,
	)
}

func ModelPath(model *string) string {
	return filepath.Join(ConfigPath(), "models", *model)
}
