// Package orb selects the default latitude/longitude lookup provider for
// standort.
//
// The provider returned by `NewProvider` uses the embedded GeoJSON asset to
// resolve geographic points to an ISO-3166 alpha-2 country code and continent
// name. The domain location service converts that continent name to the public
// two-letter continent code.
package orb
