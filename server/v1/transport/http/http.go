package http

import (
	"github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/standort/location"
)

// Register for HTTP.
func Register(location *location.Location) {
	http.Handle("/v1/ip", &ipHandler{location: location})
	http.Handle("/v1/coordinate", &coordinateHandler{location: location})
}
