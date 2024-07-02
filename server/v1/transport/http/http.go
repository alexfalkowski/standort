package http

import (
	"net/http"

	nh "github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/standort/location"
)

// Register for HTTP.
func Register(location *location.Location) {
	nh.Handle("/v1/ip", &ipHandler{location: location})
	nh.Handle("/v1/coordinate", &coordinateHandler{location: location})
}

func handleError(err error) error {
	if location.IsNotFound(err) {
		return nh.Error(http.StatusNotFound, err.Error())
	}

	return err
}
