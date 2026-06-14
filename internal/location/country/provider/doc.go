// Package provider defines the country lookup contract used by the location
// domain service.
//
// Implementations resolve ISO country codes into the normalized country code
// and continent name used as intermediate provider data. Callers should not
// expose the returned continent name directly; `internal/location` converts it
// to the public two-letter continent code.
package provider
