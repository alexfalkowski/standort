package http

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/net/http"
	"github.com/alexfalkowski/go-service/v2/net/http/status"
	v1 "github.com/alexfalkowski/standort/v2/api/standort/v1"
)

// GetLocationByIP resolves a location from an IP address.
//
// The handler delegates response construction to the v1 locator. If the
// underlying lookup returns an error, no response body is returned and the error
// is mapped to an HTTP 404 response.
func (s *Server) GetLocationByIP(ctx context.Context, req *v1.GetLocationByIPRequest) (*v1.GetLocationByIPResponse, error) {
	resp, err := s.locator.LocateByIP(ctx, req)
	if err != nil {
		return nil, status.SafeError(http.StatusNotFound, err)
	}

	return resp, nil
}

// GetLocationByLatLng resolves a location from a latitude/longitude coordinate.
//
// The handler delegates response construction to the v1 locator. If the
// underlying lookup returns an error, no response body is returned and the error
// is mapped to an HTTP 404 response.
func (s *Server) GetLocationByLatLng(ctx context.Context, req *v1.GetLocationByLatLngRequest) (*v1.GetLocationByLatLngResponse, error) {
	resp, err := s.locator.LocateByLatLng(ctx, req)
	if err != nil {
		return nil, status.SafeError(http.StatusNotFound, err)
	}

	return resp, nil
}
