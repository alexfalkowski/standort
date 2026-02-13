package assets

import (
	"embed"
)

// fs contains all embedded asset files used by standort.
//
// The embedded files are:
//   - `earth.geojson`: used for point-in-polygon lookups (lat/lng → country/continent)
//   - `geoip2.mmdb`: used for IP → country code lookups
//
//go:embed earth.geojson
//go:embed geoip2.mmdb
var fs embed.FS

// NewFS returns the embedded filesystem containing standort's runtime assets.
//
// Consumers can use the returned `embed.FS` with standard `fs` APIs such as
// `fs.ReadFile` to load embedded asset files by name (for example,
// `fs.ReadFile("earth.geojson")`).
func NewFS() embed.FS {
	return fs
}
