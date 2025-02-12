package provider

import (
	"context"
)

// Provider to get a country by an IP.
type Provider interface {
	// GetByIP a country.
	GetByIP(ctx context.Context, ip string) (string, error)
}
