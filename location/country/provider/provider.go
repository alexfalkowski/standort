package provider

import (
	"context"
	"fmt"
)

// Provider to get a country and continent.
type Provider interface {
	fmt.Stringer

	// GetByName a country and continent.
	GetByName(ctx context.Context, name string) (string, string, error)
}
