package http

import (
	"net/http"

	nh "github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/standort/server/location"
)

// Register for HTTP.
func Register(service *location.Locator) {
	nh.Handle("/v2/location", &locationHandler{service: service})
}

func handleError(err error) error {
	if location.IsNotFound(err) {
		return nh.Error(http.StatusNotFound, err.Error())
	}

	return err
}
