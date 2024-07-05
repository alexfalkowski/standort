package http

import (
	"net/http"

	nh "github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/standort/location"
)

// Register for HTTP.
func Register(location *location.Location) {
	rpc.Handle("/v1/ip", &ipHandler{location: location})
	rpc.Handle("/v1/coordinate", &coordinateHandler{location: location})
}

func handleError(err error) error {
	if location.IsNotFound(err) {
		return nh.Error(http.StatusNotFound, err.Error())
	}

	return err
}
