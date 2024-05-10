package rtree

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
)

// Node for the RTree.
type Node struct {
	Geometry  orb.Geometry
	Country   string
	Continent string
}

// IsPointInGeometry for data.
func (n *Node) IsPointInGeometry(lat, lng float64) bool {
	point := orb.Point{lng, lat}

	multiPoly, isMulti := n.Geometry.(orb.MultiPolygon)
	if isMulti {
		return planar.MultiPolygonContains(multiPoly, point)
	}

	polygon, isPoly := n.Geometry.(orb.Polygon)
	if isPoly {
		return planar.PolygonContains(polygon, point)
	}

	return false
}
