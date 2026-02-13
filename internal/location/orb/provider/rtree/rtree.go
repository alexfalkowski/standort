package rtree

import (
	"embed"
	"errors"

	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/runtime"
	"github.com/alexfalkowski/go-service/v2/strings"
	"github.com/paulmach/orb/geojson"
	"github.com/tidwall/rtree"
)

// ErrNotFound is returned when no geometry in the R-tree contains the queried point.
//
// This error is intended to be treated as a sentinel "no match" condition by
// callers, as opposed to a system failure (I/O, parsing, etc.).
var ErrNotFound = errors.New("not found")

// NewProvider constructs an R-tree-backed orb provider.
//
// It builds an in-memory spatial index from the embedded `earth.geojson` asset.
// Construction will terminate the process (via `runtime.Must`) if the GeoJSON
// asset cannot be read or parsed.
func NewProvider(fs embed.FS) *Provider {
	tree := &rtree.Generic[*Node]{}
	populateTree(tree, fs)

	return &Provider{tree: tree}
}

// Provider implements a latitude/longitude point-in-polygon search using an R-tree.
//
// The index is built from `earth.geojson`. Each node stores the geometry along
// with ISO-3166 alpha-2 country codes and continent names as provided by the dataset.
type Provider struct {
	tree *rtree.Generic[*Node]
}

// Search resolves a latitude/longitude coordinate to a country code and continent name.
//
// Inputs are in degrees. The search is performed by:
//  1. querying the R-tree by the point's bounding box to get candidate geometries, then
//  2. running an exact point-in-polygon test via `(*Node).IsPointInGeometry`.
//
// Returns:
//   - countryCode: ISO-3166 alpha-2 code from the dataset (e.g. "US")
//   - continent: continent name from the dataset (e.g. "North America")
//   - err: `ErrNotFound` when no geometry contains the point
func (p *Provider) Search(_ context.Context, lat, lng float64) (string, string, error) {
	var (
		found bool
		data  *Node
	)

	// Note: rtree uses [x,y] ordering; for geographic coordinates that is [lng,lat].
	p.tree.Search([2]float64{lng, lat}, [2]float64{lng, lat}, func(_, _ [2]float64, d *Node) bool {
		data = d

		if data.IsPointInGeometry(lat, lng) {
			found = true
			return false
		}

		return true
	})

	if !found {
		return strings.Empty, strings.Empty, ErrNotFound
	}

	return data.Country, data.Continent, nil
}

// populateTree reads `earth.geojson` from the embedded filesystem and inserts each
// feature's geometry into the R-tree.
//
// The inserted bounding boxes are derived from the feature geometry bounds.
func populateTree(tree *rtree.Generic[*Node], fs embed.FS) {
	data, err := fs.ReadFile("earth.geojson")
	runtime.Must(err)

	fc, err := geojson.UnmarshalFeatureCollection(data)
	runtime.Must(err)

	for _, f := range fc.Features {
		bound := f.Geometry.Bound()
		data := &Node{
			Country:   f.Properties["iso_a2"].(string),
			Continent: f.Properties["continent"].(string),
			Geometry:  f.Geometry,
		}

		tree.Insert(bound.Min, bound.Max, data)
	}
}
