package grpc

import (
	"context"
	"fmt"
	"net"
	"strings"

	geouri "git.jlel.se/jlelse/go-geouri"
	"github.com/alexfalkowski/go-service/meta"
	gmeta "github.com/alexfalkowski/go-service/transport/grpc/meta"
	v2 "github.com/alexfalkowski/standort/api/standort/v2"
	"github.com/alexfalkowski/standort/location"
	"google.golang.org/grpc/peer"
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
			resp.Meta = meta.Attributes(ctx)

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

	md := gmeta.ExtractIncoming(ctx)

	if f := md.Get("x-forwarded-for"); len(f) > 0 {
		return strings.Split(f[0], ",")[0]
	}

	if p, ok := peer.FromContext(ctx); ok {
		if host, _, err := net.SplitHostPort(p.Addr.String()); err != nil && host != "" {
			return host
		}
	}

	return ""
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
