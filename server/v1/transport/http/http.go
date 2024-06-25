package http

import (
	"github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/standort/location"
)

type (

	// Server for HTTP.
	Server struct {
		location *location.Location
	}

	// Error for HTTP.
	Error struct {
		Message string `json:"message,omitempty"`
	}
)

// Register for HTTP.
func Register(location *location.Location) {
	s := &Server{location: location}

	http.Handler("POST /v1/ip", &ipErrorer{}, s.GetLocationByIP)
	http.Handler("POST /v1/coordinate", &coordinateErrorer{}, s.GetLocationByLatLng)
}
