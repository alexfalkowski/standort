package http

import (
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/standort/location"
)

// Register for HTTP.
func Register(location *location.Location) {
	ih := &ipHandler{location: location}
	rpc.Route("/v1/ip", ih.Locate)

	ch := &coordinateHandler{location: location}
	rpc.Route("/v1/coordinate", ch.Locate)
}

func handleError(err error) error {
	if location.IsNotFound(err) {
		return status.Error(http.StatusNotFound, err.Error())
	}

	return err
}
