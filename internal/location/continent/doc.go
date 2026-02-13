// Package continent provides a normalization map for continent identifiers.
//
// Different upstream data sources commonly represent continents as English names
// (for example "Europe" or "North America"). standort normalizes these names to
// stable, two-letter continent codes that are returned by the domain
// `internal/location` service and, in turn, the public API.
//
// # Codes
//
// The `Codes` map contains the canonical mapping from continent name to
// two-letter code:
//
//   - "Africa"        → "AF"
//   - "Asia"          → "AS"
//   - "Europe"        → "EU"
//   - "North America" → "NA"
//   - "Oceania"       → "OC"
//   - "South America" → "SA"
//
// Callers should treat the keys as the provider-emitted continent name strings.
// If a name is not present in the map, the lookup will return the empty string,
// and callers may choose to handle that as an unknown/unsupported continent.
package continent
