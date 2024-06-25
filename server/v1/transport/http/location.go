package http

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/standort/location"
)

type (
	// GetLocationByIPRequest for an IP address.
	GetLocationByIPRequest struct {
		IP string `json:"ip,omitempty"`
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
	}

	ipHandler struct {
		location *location.Location
	}
	coordinateHandler struct {
		location *location.Location
	}
)

func (h *ipHandler) Handle(ctx context.Context, req *GetLocationByIPRequest) (*GetLocationByIPResponse, error) {
	resp := &GetLocationByIPResponse{}

	country, continent, err := h.location.GetByIP(ctx, req.IP)
	if err != nil {
		return resp, err
	}

	resp.Location = &Location{Country: country, Continent: continent}
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}

func (h *ipHandler) Error(ctx context.Context, err error) *GetLocationByIPResponse {
	return &GetLocationByIPResponse{Meta: meta.CamelStrings(ctx, ""), Error: &Error{Message: err.Error()}}
}

func (h *ipHandler) Status(err error) int {
	if location.IsNotFound(err) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

func (h *coordinateHandler) Handle(ctx context.Context, req *GetLocationByLatLngRequest) (*GetLocationByLatLngResponse, error) {
	resp := &GetLocationByLatLngResponse{Location: &Location{}}

	country, continent, err := h.location.GetByLatLng(ctx, req.Lat, req.Lng)
	if err != nil {
		return resp, err
	}

	resp.Location = &Location{Country: country, Continent: continent}
	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}

func (h *coordinateHandler) Error(ctx context.Context, err error) *GetLocationByLatLngResponse {
	return &GetLocationByLatLngResponse{Meta: meta.CamelStrings(ctx, ""), Error: &Error{Message: err.Error()}}
}

func (h *coordinateHandler) Status(err error) int {
	if location.IsNotFound(err) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
