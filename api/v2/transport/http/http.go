package http

import (
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/rpc"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/standort/api/location"
	"github.com/alexfalkowski/standort/location/errors"
)

// Register for HTTP.
func Register(handler *Handler) {
	rpc.Route("/v2/location", handler.GetLocation)
}

// NewHandler for HTTP.
func NewHandler(locator *location.Locator) *Handler {
	return &Handler{locator: locator}
}

// Handler for HTTP.
type Handler struct {
	locator *location.Locator
}

func (h *Handler) error(err error) error {
	if err == nil {
		return nil
	}

	if errors.IsNotFound(err) {
		return status.Error(http.StatusNotFound, err.Error())
	}

	return err
}
