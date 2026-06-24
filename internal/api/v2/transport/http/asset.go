package http

import (
	"github.com/alexfalkowski/go-service/v2/context"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
)

// GetLookupAssets returns embedded lookup asset metadata.
func (s *Server) GetLookupAssets(ctx context.Context, _ *v2.GetLookupAssetsRequest) (*v2.GetLookupAssetsResponse, error) {
	return s.assets.Get(ctx), nil
}
