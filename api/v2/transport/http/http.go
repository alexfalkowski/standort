package http

import (
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/standort/api/location"
)

// Register for HTTP.
func Register(service *location.Locator) {
	lh := &locationHandler{service: service}
	rpc.Route("/v2/location", lh.Locate)
}

func handleError(err error) error {
	if location.IsNotFound(err) {
		return status.Error(http.StatusNotFound, err.Error())
	}

	return err
}
