package http

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/net/http"
	"github.com/alexfalkowski/go-service/v2/net/http/status"
	v1 "github.com/alexfalkowski/standort/v2/api/standort/v1"
	"github.com/alexfalkowski/standort/v2/internal/api/v1/location"
)

func getLocationByIP(locator *location.Locator) func(context.Context, *v1.GetLocationByIPRequest) (*v1.GetLocationByIPResponse, error) {
	return func(ctx context.Context, req *v1.GetLocationByIPRequest) (*v1.GetLocationByIPResponse, error) {
		resp, err := locator.LocateByIP(ctx, req)
		if err != nil {
			return nil, status.SafeError(http.StatusNotFound, err)
		}

		return resp, nil
	}
}

func getLocationByLatLng(locator *location.Locator) func(context.Context, *v1.GetLocationByLatLngRequest) (*v1.GetLocationByLatLngResponse, error) {
	return func(ctx context.Context, req *v1.GetLocationByLatLngRequest) (*v1.GetLocationByLatLngResponse, error) {
		resp, err := locator.LocateByLatLng(ctx, req)
		if err != nil {
			return nil, status.SafeError(http.StatusNotFound, err)
		}

		return resp, nil
	}
}
