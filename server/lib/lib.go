package lib

import (
	"os"
	"path/filepath"
)

func ModelPath(model *string) string {
    if model == nil {
        return filepath.Join(os.Getenv("HOME"), ".portal", "models")
    } else {
        return filepath.Join(os.Getenv("HOME"), ".portal", "models", *model)
    }
}
