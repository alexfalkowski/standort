package grpc

import (
	"context"
	"fmt"

	geouri "git.jlel.se/jlelse/go-geouri"
	"github.com/alexfalkowski/go-service/meta"
	gmeta "github.com/alexfalkowski/go-service/transport/grpc/meta"
	tmeta "github.com/alexfalkowski/go-service/transport/meta"
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
			meta.WithAttribute(ctx, "location.ip_error", err.Error())
		} else {
			resp.Locations = append(resp.Locations, &v2.Location{Country: country, Continent: continent, Kind: v2.Kind_KIND_IP})
		}
	}

	point, err := s.point(ctx, req)
	if err != nil {
		meta.WithAttribute(ctx, "location.point_error", err.Error())
	} else {
		if point == nil {
			return resp, nil
		}

		if country, continent, err := s.location.GetByLatLng(ctx, point.Lat, point.Lng); err != nil {
			meta.WithAttribute(ctx, "location.lat_lng_error", err.Error())
		} else {
			resp.Locations = append(resp.Locations, &v2.Location{Country: country, Continent: continent, Kind: v2.Kind_KIND_GEO})
		}
	}

	resp.Meta = meta.Attributes(ctx)

	return resp, nil
}

func (s *Server) ip(ctx context.Context, req *v2.GetLocationRequest) string {
	if req.Ip != "" {
		return req.Ip
	}

	return tmeta.RemoteAddress(ctx)
}

func (s *Server) point(ctx context.Context, req *v2.GetLocationRequest) (*v2.Point, error) {
	if req.Point != nil {
		return req.Point, nil
	}

	md := gmeta.ExtractIncoming(ctx)

	values := md.Get("geolocation")
	if len(values) > 0 {
		geo, err := geouri.Parse(values[0])
		if err != nil {
			return nil, fmt.Errorf("geo uri: %w", err)
		}

		return &v2.Point{Lat: geo.Latitude, Lng: geo.Longitude}, nil
	}

	return nil, nil
}
