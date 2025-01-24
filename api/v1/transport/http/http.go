package http

import (
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/standort/location"
)

// Register for HTTP.
func Register(handler *Handler) {
	rpc.Route("/v1/ip", handler.GetLocationByIP)
	rpc.Route("/v1/coordinate", handler.GetLocationByLatLng)
}

// NewHandler for HTTP.
func NewHandler(location *location.Location) *Handler {
	return &Handler{location: location}
}

// Handler for HTTP.
type Handler struct {
	location *location.Location
}

func (h *Handler) error(err error) error {
	if err == nil {
		return nil
	}

	return status.Error(http.StatusNotFound, err.Error())
}
