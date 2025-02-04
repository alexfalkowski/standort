package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/types/slices"
	"github.com/alexfalkowski/standort/api/location"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
)

// Kinds maps from location.Kind to v2.Kind.
var Kinds = map[location.Kind]v2.Kind{
	location.IP:  v2.Kind_KIND_IP,
	location.GEO: v2.Kind_KIND_GEO,
}

// GetLocation for gRPC.
func (s *Server) GetLocation(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
	resp := &v2.GetLocationResponse{}
	locations := []*v2.Location{}
	ip, geo, err := s.service.Locate(ctx, req.GetIp(), toPoint(req.GetPoint()))
	locations = slices.AppendNotNil(locations, toLocation(ip), toLocation(geo))

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.Locations = locations

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

	return &v2.Location{Country: l.Country, Continent: l.Continent, Kind: Kinds[l.Kind]}
}
