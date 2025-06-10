package assets

import (
	"embed"
)

//go:embed earth.geojson
//go:embed geoip2.mmdb
var fs embed.FS

// NewFS for assets.
func NewFS() embed.FS {
	return fs
}
