package http

import (
	"github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/standort/location"
)

// Error for HTTP.
type Error struct {
	Message string `json:"message,omitempty"`
}

// Register for HTTP.
func Register(location *location.Location) {
	http.Handle("POST /v1/ip", &ipHandler{location: location})
	http.Handle("POST /v1/coordinate", &coordinateHandler{location: location})
}
