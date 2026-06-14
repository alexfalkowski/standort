// Package rtree implements point-in-polygon lookup with an in-memory R-tree.
//
// The provider builds its index from the embedded `earth.geojson` asset. Queries
// first use the R-tree to find candidate geometries by bounding box, then run an
// exact point-in-polygon check before returning the dataset's country code and
// continent name.
package rtree
