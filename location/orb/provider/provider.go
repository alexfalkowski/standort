package provider

import (
	"context"
	"fmt"
)

// Provider to search a lat lng and get country and continent.
type Provider interface {
	fmt.Stringer

	// Search a lat lng and get country and continent.
	Search(ctx context.Context, lat, lng float64) (string, string)
}
