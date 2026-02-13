package rtree

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
)

// Node is the value stored in the R-tree index.
//
// Each node represents a geographic region (a polygon or multipolygon) and the
// associated location identifiers derived from the GeoJSON dataset:
//
//   - Country: ISO-3166 alpha-2 country code (e.g. "US").
//   - Continent: continent name (e.g. "North America").
//
// Geometry is expected to be either an `orb.Polygon` or an `orb.MultiPolygon`.
type Node struct {
	Geometry  orb.Geometry
	Country   string
	Continent string
}

// IsPointInGeometry reports whether the provided latitude/longitude point lies
// within this node's geometry.
//
// Inputs are expressed in degrees. Note that orb uses [x,y] ordering, which for
// geographic coordinates corresponds to [lng,lat].
func (n *Node) IsPointInGeometry(lat, lng float64) bool {
	point := orb.Point{lng, lat}

	multiPoly, ok := n.Geometry.(orb.MultiPolygon)
	if ok {
		return planar.MultiPolygonContains(multiPoly, point)
	}

	return planar.PolygonContains(n.Geometry.(orb.Polygon), point)
}
