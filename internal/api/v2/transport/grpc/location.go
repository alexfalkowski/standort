package grpc

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/location"
)

// GetLocation resolves a location using any available inputs (IP address and/or geo point).
//
// This handler delegates to `(*location.Locator).Locate`, which applies transport-level
// behavior such as:
//   - reading fallback inputs from request metadata (if not present in the request), and
//   - attaching lookup/parsing failures to metadata attributes.
//
// Response semantics:
//   - `resp.Meta` is populated from request metadata via `meta.CamelStrings(ctx, strings.Empty)`.
//   - `resp.Ip` is set when an IP-derived location could be resolved.
//   - `resp.Geo` is set when a geo-derived location could be resolved.
//   - if neither input produces a location, the resulting error is mapped to a gRPC
//     `codes.NotFound` status by `s.error`.
//
// On partial success (one lookup succeeds and the other fails or is missing), the
// successful side is returned and the other side is nil.
func (s *Server) GetLocation(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
	ip, geo, err := s.service.Locate(ctx, req.GetIp(), toPoint(req.GetPoint()))
	resp := &v2.GetLocationResponse{
		Meta: meta.CamelStrings(ctx, strings.Empty),
		Ip:   toLocation(ip),
		Geo:  toLocation(geo),
	}

	return resp, s.error(err)
}

// toPoint converts the generated protobuf `v2.Point` into the transport-facing
// `location.Point` type used by the service layer.
//
// A nil input returns nil.
func toPoint(p *v2.Point) *location.Point {
	if p == nil {
		return nil
	}

	return &location.Point{Lat: p.GetLat(), Lng: p.GetLng()}
}

// toLocation converts the transport-facing `location.Location` returned by the
// service into the generated protobuf `v2.Location` message.
//
// A nil input returns nil.
func toLocation(l *location.Location) *v2.Location {
	if l == nil {
		return nil
	}

	return &v2.Location{Country: l.Country, Continent: l.Continent}
}
