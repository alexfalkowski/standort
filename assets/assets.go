package assets

import (
	"embed"
)

//go:embed *.geojson
//go:embed *.mmdb
//go:embed *.bin
var fs embed.FS

// FS for assets.
func FS() embed.FS {
	return fs
}
