package assets

import (
	"embed"
	"io/fs"
)

//go:embed *
var assets embed.FS

func FS() fs.FS {
	return assets
}

func Asset(name string) ([]byte, error) {
	return assets.ReadFile(name)
}
