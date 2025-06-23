package provider

import "github.com/alexfalkowski/go-service/v2/context"

// Provider to get a country and continent.
type Provider interface {
	// GetByCode a country and continent.
	GetByCode(ctx context.Context, code string) (string, string, error)
}
