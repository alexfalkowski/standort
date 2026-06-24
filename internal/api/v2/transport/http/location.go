package http

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/net/http"
	"github.com/alexfalkowski/go-service/v2/net/http/status"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/location"
	"github.com/alexfalkowski/standort/v2/internal/diagnostics"
)

func getLocation(locator *location.Locator) func(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
	return func(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
		resp, err := locator.Locate(ctx, req)
		if err != nil {
			setFailureHeaders(ctx, diagnostics.FromError(err))
			return nil, status.SafeError(http.StatusNotFound, err)
		}

		return resp, nil
	}
}
