package http

import (
	"github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/standort/server/service"
)

// Register for HTTP.
func Register(service *service.Service) {
	http.Handle("/v2/location", &locationHandler{service: service})
}
