package http

import (
	"context"
	"fmt"
	"net/http"

	geouri "git.jlel.se/jlelse/go-geouri"
	"github.com/alexfalkowski/go-service/meta"
	tm "github.com/alexfalkowski/go-service/transport/meta"
	"github.com/alexfalkowski/standort/location"
)

type (
	// GetLocationRequest for getting the location.
	GetLocationRequest struct {
		Point *Point `json:"point,omitempty"`
		IP    string `json:"ip,omitempty"`
	}

	// GetLocationResponse for getting the location.
	GetLocationResponse struct {
		Meta      map[string]string `json:"meta,omitempty"`
		Error     *Error            `json:"error,omitempty"`
		Locations []*Location       `json:"locations,omitempty"`
	}

	// Point for the request.
	Point struct {
		Lat float64 `json:"lat,omitempty"`
		Lng float64 `json:"lng,omitempty"`
	}

	// GetLocationByIPResponse for an IP address.
	GetLocationByIPResponse struct {
		Meta     map[string]string `json:"meta,omitempty"`
		Error    *Error            `json:"error,omitempty"`
		Location *Location         `json:"location,omitempty"`
	}

	// GetLocationByLatLngRequest for a latitude and longitude.
	GetLocationByLatLngRequest struct {
		Lat float64 `json:"lat,omitempty"`
		Lng float64 `json:"lng,omitempty"`
	}

	// GetLocationByLatLngResponse for a latitude and longitude.
	GetLocationByLatLngResponse struct {
		Meta     map[string]string `json:"meta,omitempty"`
		Error    *Error            `json:"error,omitempty"`
		Location *Location         `json:"location,omitempty"`
	}

	// Location of the response.
	Location struct {
		Country   string `json:"country,omitempty"`
		Continent string `json:"continent,omitempty"`
		Kind      string `json:"kind,omitempty"`
	}

	locationHandler struct {
		location *location.Location
	}
)

func (h *locationHandler) Handle(ctx context.Context, req *GetLocationRequest) (*GetLocationResponse, error) {
	resp := &GetLocationResponse{Locations: []*Location{}}

	if ip := h.ip(ctx, req); ip != "" {
		if country, continent, err := h.location.GetByIP(ctx, ip); err != nil {
			meta.WithAttribute(ctx, "location.ip_error", meta.Error(err))
		} else {
			resp.Locations = append(resp.Locations, &Location{Country: country, Continent: continent, Kind: "KIND_IP"})
		}
	}

	point, err := h.point(ctx, req)
	if err != nil {
		meta.WithAttribute(ctx, "location.point_error", meta.Error(err))
	} else {
		if point == nil {
			resp.Meta = meta.CamelStrings(ctx, "")

			return resp, nil
		}

		if country, continent, err := h.location.GetByLatLng(ctx, point.Lat, point.Lng); err != nil {
			meta.WithAttribute(ctx, "location.lat_lng_error", meta.Error(err))
		} else {
			resp.Locations = append(resp.Locations, &Location{Country: country, Continent: continent, Kind: "KIND_GEO"})
		}
	}

	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}

func (h *locationHandler) Error(ctx context.Context, err error) *GetLocationResponse {
	return &GetLocationResponse{Meta: meta.CamelStrings(ctx, ""), Error: &Error{Message: err.Error()}}
}

func (h *locationHandler) Status(err error) int {
	if location.IsNotFound(err) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

func (h *locationHandler) ip(ctx context.Context, req *GetLocationRequest) string {
	ip := req.IP
	if ip != "" {
		return ip
	}

	return tm.IPAddr(ctx).Value()
}

func (h *locationHandler) point(ctx context.Context, req *GetLocationRequest) (*Point, error) {
	point := req.Point
	if point != nil {
		return point, nil
	}

	l := tm.Geolocation(ctx).Value()
	if l == "" {
		return nil, nil //nolint:nilnil
	}

	geo, err := geouri.Parse(l)
	if err != nil {
		return nil, fmt.Errorf("geo uri: %w", err)
	}

	return &Point{Lat: geo.Latitude, Lng: geo.Longitude}, nil
}
