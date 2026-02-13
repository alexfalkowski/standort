package provider

import "github.com/alexfalkowski/go-service/v2/context"

// Provider resolves an IP address to a country code.
//
// Implementations are expected to return an ISO-3166 alpha-2 country code
// (for example, "US") when a match exists.
type Provider interface {
	// GetByIP resolves the ISO-3166 alpha-2 country code for the given IP address.
	//
	// The `ip` string should be a textual representation of an IPv4 or IPv6 address.
	//
	// Returns:
	//   - countryCode: ISO-3166 alpha-2 code (e.g. "US")
	//   - err: any lookup/validation error
	GetByIP(ctx context.Context, ip string) (countryCode string, err error)
}
