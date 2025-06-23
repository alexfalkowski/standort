package rtree

import (
	"embed"
	"errors"

	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/runtime"
	"github.com/paulmach/orb/geojson"
	"github.com/tidwall/rtree"
)

// ErrNotFound for rtree.
var ErrNotFound = errors.New("not found")

// NewProvider for rtree.
func NewProvider(fs embed.FS) *Provider {
	tree := &rtree.Generic[*Node]{}
	populateTree(tree, fs)

	return &Provider{tree: tree}
}

// Provider for rtree.
type Provider struct {
	tree *rtree.Generic[*Node]
}

// Search a lat lng and get country and continent.
func (p *Provider) Search(_ context.Context, lat, lng float64) (string, string, error) {
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
		return "", "", ErrNotFound
	}

	return data.Country, data.Continent, nil
}

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
