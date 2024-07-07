package http

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/standort/server/location"
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
		Location *Location         `json:"location,omitempty"`
	}

	// Location of the response.
	Location struct {
		Country   string `json:"country,omitempty"`
		Continent string `json:"continent,omitempty"`
		Kind      string `json:"kind,omitempty"`
	}

	locationHandler struct {
		service *location.Locator
	}
)

func (h *locationHandler) Handle(ctx context.Context, req *GetLocationRequest) (*GetLocationResponse, error) {
	resp := &GetLocationResponse{}
	locations := []*Location{}

	ip, geo, err := h.service.Locate(ctx, req.IP, toPoint(req.Point))
	if err != nil {
		resp.Meta = meta.CamelStrings(ctx, "")

		return resp, handleError(err)
	}

	i, g := toLocation(ip), toLocation(geo)

	if i != nil {
		locations = append(locations, i)
	}

	if g != nil {
		locations = append(locations, g)
	}

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.Locations = locations

	return resp, nil
}

func toPoint(p *Point) *location.Point {
	if p == nil {
		return nil
	}

	return &location.Point{Lat: p.Lat, Lng: p.Lng}
}

func toLocation(l *location.Location) *Location {
	if l == nil {
		return nil
	}

	return &Location{Country: l.Country, Continent: l.Continent, Kind: string(l.Kind)}
}
