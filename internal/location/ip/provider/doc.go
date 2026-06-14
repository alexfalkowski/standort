// Package provider defines the IP lookup contract used by the location domain
// service.
//
// Implementations resolve textual IPv4 or IPv6 addresses to ISO-3166 alpha-2
// country codes. Country-to-continent enrichment happens later through the
// country provider owned by `internal/location`.
package provider
