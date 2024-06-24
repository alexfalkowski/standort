package grpc

import (
	"context"
	"fmt"

	geouri "git.jlel.se/jlelse/go-geouri"
	"github.com/alexfalkowski/go-service/meta"
	tm "github.com/alexfalkowski/go-service/transport/meta"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
)

// GetLocation for gRPC.
func (s *Server) GetLocation(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
	resp := &v2.GetLocationResponse{Locations: []*v2.Location{}}

	if ip := s.ip(ctx, req); ip != "" {
		if country, continent, err := s.location.GetByIP(ctx, ip); err != nil {
			meta.WithAttribute(ctx, "location.ip_error", meta.Error(err))
		} else {
			resp.Locations = append(resp.GetLocations(), &v2.Location{Country: country, Continent: continent, Kind: v2.Kind_KIND_IP})
		}
	}

	point, err := s.point(ctx, req)
	if err != nil {
		meta.WithAttribute(ctx, "location.point_error", meta.Error(err))
	} else {
		if point == nil {
			resp.Meta = meta.CamelStrings(ctx, "")

			return resp, nil
		}

		if country, continent, err := s.location.GetByLatLng(ctx, point.GetLat(), point.GetLng()); err != nil {
			meta.WithAttribute(ctx, "location.lat_lng_error", meta.Error(err))
		} else {
			resp.Locations = append(resp.GetLocations(), &v2.Location{Country: country, Continent: continent, Kind: v2.Kind_KIND_GEO})
		}
	}

	resp.Meta = meta.CamelStrings(ctx, "")

	return resp, nil
}

func (s *Server) ip(ctx context.Context, req *v2.GetLocationRequest) string {
	ip := req.GetIp()
	if ip != "" {
		return ip
	}

	return tm.IPAddr(ctx).Value()
}

func (s *Server) point(ctx context.Context, req *v2.GetLocationRequest) (*v2.Point, error) {
	point := req.GetPoint()
	if point != nil {
		return point, nil
	}

	l := tm.Geolocation(ctx).Value()
	if l == "" {
		return nil, nil //nolint:nilnil
	}

	geo, err := geouri.Parse(l)
	if err != nil {
		return nil, fmt.Errorf("geo uri: %w", err)
	}

	return &v2.Point{Lat: geo.Latitude, Lng: geo.Longitude}, nil
}
