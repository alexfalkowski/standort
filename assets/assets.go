package assets

import (
	"embed"
)

//go:embed earth.geojson
//go:embed geoip2.mmdb
var fs embed.FS

// FS for assets.
func FS() embed.FS {
	return fs
}
