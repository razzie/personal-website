package assets

import (
	"embed"
	"io"
	"io/fs"
)

//go:embed *
var assets embed.FS

func FS() fs.FS {
	return assets
}

func Asset(name string) ([]byte, error) {
	file, err := assets.Open(name)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(file)
}
