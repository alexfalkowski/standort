package provider

import "github.com/alexfalkowski/go-service/v2/context"

// Provider to search a lat lng and get country and continent.
type Provider interface {
	// Search a lat lng and get country and continent.
	Search(ctx context.Context, lat, lng float64) (string, string, error)
}
