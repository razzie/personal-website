package assets

//go:generate go run ../../tools/go-bindata/ -pkg assets -prefix ../../assets ../../assets/...

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
)

// FS ...
func FS(prefix string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: nil,
		Prefix:    prefix,
	}
}
