package grpc

import (
	"context"
	"fmt"

	geouri "git.jlel.se/jlelse/go-geouri"
	"github.com/alexfalkowski/go-service/meta"
	sm "github.com/alexfalkowski/go-service/transport/grpc/meta"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	"github.com/alexfalkowski/standort/location"
)

// NewServer for gRPC.
func NewServer(location *location.Location) v2.ServiceServer {
	return &Server{location: location}
}

// Server for gRPC.
type Server struct {
	location *location.Location
	v2.UnimplementedServiceServer
}

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
			resp.Meta = s.meta(ctx)

			return resp, nil
		}

		if country, continent, err := s.location.GetByLatLng(ctx, point.GetLat(), point.GetLng()); err != nil {
			meta.WithAttribute(ctx, "location.lat_lng_error", meta.Error(err))
		} else {
			resp.Locations = append(resp.GetLocations(), &v2.Location{Country: country, Continent: continent, Kind: v2.Kind_KIND_GEO})
		}
	}

	resp.Meta = s.meta(ctx)

	return resp, nil
}

func (s *Server) ip(ctx context.Context, req *v2.GetLocationRequest) string {
	ip := req.GetIp()
	if ip != "" {
		return ip
	}

	return sm.IPAddr(ctx, sm.ExtractIncoming(ctx))
}

func (s *Server) point(ctx context.Context, req *v2.GetLocationRequest) (*v2.Point, error) {
	point := req.GetPoint()
	if point != nil {
		return point, nil
	}

	md := sm.ExtractIncoming(ctx)

	values := md.Get("geolocation")
	if len(values) > 0 {
		geo, err := geouri.Parse(values[0])
		if err != nil {
			return nil, fmt.Errorf("geo uri: %w", err)
		}

		return &v2.Point{Lat: geo.Latitude, Lng: geo.Longitude}, nil
	}

	return nil, nil //nolint:nilnil
}

func (s *Server) meta(ctx context.Context) map[string]string {
	return meta.CamelStrings(ctx, "")
}
