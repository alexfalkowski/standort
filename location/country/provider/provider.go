package provider

import (
	"context"
)

// Provider to get a country and continent.
type Provider interface {
	// GetByName a country and continent.
	GetByName(ctx context.Context, name string) (string, string, error)
}
