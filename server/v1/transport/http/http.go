package http

import (
	"github.com/alexfalkowski/go-service/marshaller"
	"github.com/alexfalkowski/go-service/net/http"
	"github.com/alexfalkowski/standort/location"
	"go.uber.org/fx"
)

type (
	// RegisterParams for HTTP.
	RegisterParams struct {
		fx.In

		Marshaller *marshaller.Map
		Mux        http.ServeMux
		Location   *location.Location
	}

	// Server for HTTP.
	Server struct {
		location *location.Location
	}

	// Error for HTTP.
	Error struct {
		Message string `json:"message,omitempty"`
	}
)

// Register for HTTP.
func Register(params RegisterParams) error {
	s := &Server{location: params.Location}

	ih := http.NewHandler[GetLocationByIPRequest](params.Mux, params.Marshaller, &ipErrorer{})
	if err := ih.Handle("POST", "/v1/ip", s.GetLocationByIP); err != nil {
		return err
	}

	ch := http.NewHandler[GetLocationByLatLngRequest](params.Mux, params.Marshaller, &coordinateErrorer{})
	if err := ch.Handle("POST", "/v1/coordinate", s.GetLocationByLatLng); err != nil {
		return err
	}

	return nil
}
