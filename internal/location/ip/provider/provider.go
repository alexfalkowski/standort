package provider

import "github.com/alexfalkowski/go-service/v2/context"

// Provider to get a country by an IP.
type Provider interface {
	// GetByIP a country.
	GetByIP(ctx context.Context, ip string) (string, error)
}
