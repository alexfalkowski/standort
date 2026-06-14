// Package country selects the default country lookup provider for standort.
//
// The provider returned by `NewProvider` resolves an ISO country code to the
// normalized country code and continent name expected by the domain
// `internal/location` service. The domain service maps that continent name to
// the public two-letter continent code.
package country
