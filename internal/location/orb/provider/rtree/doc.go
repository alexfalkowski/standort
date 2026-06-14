// Package rtree implements point-in-polygon lookup with an in-memory R-tree.
//
// The provider builds its index from the embedded `earth.geojson` asset. Queries
// first use the R-tree to find candidate geometries by bounding box, then run an
// exact point-in-polygon check before returning the dataset's country code and
// continent name.
//
// Only supported GeoJSON features are indexed. A feature must provide a
// two-character `iso_a2` or `iso_a2_eh` country code and a `continent` value
// known to `internal/location/continent.Codes`. Other features are skipped, so a
// not-found result means no indexed supported geometry matched the point.
package rtree
