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

	multiPoly, ok := n.Geometry.(orb.MultiPolygon)
	if ok {
		return planar.MultiPolygonContains(multiPoly, point)
	}

	return planar.PolygonContains(n.Geometry.(orb.Polygon), point)
}
