package provider

import (
	"context"
	"fmt"
)

// Provider to get a country by an IP.
type Provider interface {
	fmt.Stringer

	// GetByIP a country.
	GetByIP(ctx context.Context, ip string) (string, error)
}
