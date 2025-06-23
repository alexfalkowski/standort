package grpc

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/location"
)

// GetLocation for gRPC.
func (s *Server) GetLocation(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
	ip, geo, err := s.service.Locate(ctx, req.GetIp(), toPoint(req.GetPoint()))
	resp := &v2.GetLocationResponse{
		Meta: meta.CamelStrings(ctx, ""),
		Ip:   toLocation(ip),
		Geo:  toLocation(geo),
	}

	return resp, s.error(err)
}

func toPoint(p *v2.Point) *location.Point {
	if p == nil {
		return nil
	}

	return &location.Point{Lat: p.GetLat(), Lng: p.GetLng()}
}

func toLocation(l *location.Location) *v2.Location {
	if l == nil {
		return nil
	}

	return &v2.Location{Country: l.Country, Continent: l.Continent}
}
