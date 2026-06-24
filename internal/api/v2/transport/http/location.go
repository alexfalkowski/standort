package http

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/net/http"
	"github.com/alexfalkowski/go-service/v2/net/http/status"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/diagnostics"
)

// GetLocation resolves a location from request inputs.
//
// The handler delegates response construction to the v2 locator. If the
// underlying lookup returns an error, diagnostic values are written as HTTP
// headers, no response body is returned, and the error is mapped to an HTTP 404
// response.
func (s *Server) GetLocation(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
	resp, err := s.locator.Locate(ctx, req)
	if err != nil {
		setFailureHeaders(ctx, diagnostics.FromError(err))
		return nil, status.SafeError(http.StatusNotFound, err)
	}

	return resp, nil
}
