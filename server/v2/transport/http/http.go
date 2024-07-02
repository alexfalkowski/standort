package http

import (
	"net/http"

	nh "github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/standort/location"
	"github.com/alexfalkowski/standort/server/service"
)

// Register for HTTP.
func Register(service *service.Service) {
	nh.Handle("/v2/location", &locationHandler{service: service})
}

func handleError(err error) error {
	if service.IsNotFound(err) || location.IsNotFound(err) {
		return nh.Error(http.StatusNotFound, err.Error())
	}

	return err
}
