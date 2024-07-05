package http

import (
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/standort/server/location"
)

// Register for HTTP.
func Register(service *location.Locator) {
	rpc.Handle("/v2/location", &locationHandler{service: service})
}

func handleError(err error) error {
	if location.IsNotFound(err) {
		return rpc.Error(http.StatusNotFound, err.Error())
	}

	return err
}
