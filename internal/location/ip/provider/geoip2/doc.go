// Package geoip2 adapts the embedded GeoIP2 country database to standort's IP
// provider contract.
//
// The provider loads `geoip2.mmdb` from the supplied embedded filesystem during
// construction and resolves IP addresses to ISO-3166 alpha-2 country codes.
package geoip2
