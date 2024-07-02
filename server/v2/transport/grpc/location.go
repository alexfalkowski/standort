package grpc

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	"github.com/alexfalkowski/standort/server/service"
)

// GetLocation for gRPC.
func (s *Server) GetLocation(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
	resp := &v2.GetLocationResponse{}
	locations := []*v2.Location{}

	ip, geo, err := s.service.GetLocations(ctx, req.GetIp(), toPoint(req.GetPoint()))
	if err != nil {
		resp.Meta = meta.CamelStrings(ctx, "")

		return resp, s.error(err)
	}

	i, g := toLocation(ip), toLocation(geo)

	if i != nil {
		locations = append(locations, i)
	}

	if g != nil {
		locations = append(locations, g)
	}

	resp.Meta = meta.CamelStrings(ctx, "")
	resp.Locations = locations

	return resp, nil
}

func toPoint(p *v2.Point) *service.Point {
	if p == nil {
		return nil
	}

	return &service.Point{Lat: p.GetLat(), Lng: p.GetLng()}
}

func toLocation(l *service.Location) *v2.Location {
	if l == nil {
		return nil
	}

	var k v2.Kind

	switch l.Kind {
	case service.GEO:
		k = v2.Kind_KIND_GEO
	case service.IP:
		k = v2.Kind_KIND_IP
	default:
		k = v2.Kind_KIND_UNSPECIFIED
	}

	return &v2.Location{Country: l.Country, Continent: l.Continent, Kind: k}
}
