package rtree

import (
	"context"
	"os"
	"path/filepath"

	"github.com/alexfalkowski/standort/location/continent"
	"github.com/paulmach/orb/geojson"
	"github.com/tidwall/rtree"
)

// Provider for rtree.
type Provider struct {
	tree *rtree.Generic[*Node]
}

// NewProvider for rtree.
func NewProvider(cfg *continent.Config) (*Provider, error) {
	paths := []string{
		cfg.GetAfricaPath(), cfg.GetAsiaPath(), cfg.GetEuropePath(),
		cfg.GetNorthAmericaPath(), cfg.GetOceaniaPath(), cfg.GetSouthAmericaPath(),
	}
	tree := &rtree.Generic[*Node]{}

	for _, path := range paths {
		if err := populateTree(tree, path); err != nil {
			return nil, err
		}
	}

	return &Provider{tree: tree}, nil
}

// Search a lat lng and get country and continent.
func (p *Provider) Search(ctx context.Context, lat, lng float64) (string, string) {
	var (
		found bool
		data  *Node
	)

	p.tree.Search([2]float64{lng, lat}, [2]float64{lng, lat}, func(min, max [2]float64, d *Node) bool {
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

func populateTree(tree *rtree.Generic[*Node], path string) error {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return err
	}

	fc, err := geojson.UnmarshalFeatureCollection(data)
	if err != nil {
		return err
	}

	for _, f := range fc.Features {
		bound := f.Geometry.Bound()
		data := &Node{
			Country:   f.Properties["iso_a2"].(string),
			Continent: f.Properties["continent"].(string),
			Geometry:  f.Geometry,
		}

		tree.Insert(bound.Min, bound.Max, data)
	}

	return nil
}
