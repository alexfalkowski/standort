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

	lh := http.NewHandler[GetLocationRequest](params.Mux, params.Marshaller, &locationErrorer{})
	if err := lh.Handle("POST", "/v2/location", s.GetLocation); err != nil {
		return err
	}

	return nil
}
