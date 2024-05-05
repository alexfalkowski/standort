package rtree

import (
	"context"
	"embed"

	"github.com/alexfalkowski/go-service/runtime"
	"github.com/paulmach/orb/geojson"
	"github.com/tidwall/rtree"
)

// Provider for rtree.
type Provider struct {
	tree *rtree.Generic[*Node]
}

// NewProvider for rtree.
func NewProvider(fs embed.FS) *Provider {
	paths := []string{
		"africa.geojson", "north_america.geojson", "oceania.geojson",
		"asia.geojson", "europe.geojson", "south_america.geojson",
	}
	tree := &rtree.Generic[*Node]{}

	for _, path := range paths {
		populateTree(tree, fs, path)
	}

	return &Provider{tree: tree}
}

// Search a lat lng and get country and continent.
func (p *Provider) Search(_ context.Context, lat, lng float64) (string, string) {
	var (
		found bool
		data  *Node
	)

	p.tree.Search([2]float64{lng, lat}, [2]float64{lng, lat}, func(_, _ [2]float64, d *Node) bool {
		data = d

		if data.IsPointInGeometry(lat, lng) {
			found = true

			return false
		}

		return true
	})

	if !found {
		return "", ""
	}

	return data.Country, data.Continent
}

func populateTree(tree *rtree.Generic[*Node], fs embed.FS, path string) {
	data, err := fs.ReadFile(path)
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
