package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/alexfalkowski/go-service/meta"
	v1 "github.com/alexfalkowski/standort/api/standort/v1"
	"github.com/alexfalkowski/standort/location/continent"
	"github.com/ip2location/ip2location-go/v9"
	"github.com/pariz/gountries"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server for gRPC.
type Server struct {
	db    *ip2location.DB
	query *gountries.Query
	v1.UnimplementedServiceServer
}

// GetLocationByIP for gRPC.
func (s *Server) GetLocationByIP(ctx context.Context, req *v1.GetLocationByIPRequest) (*v1.GetLocationByIPResponse, error) {
	resp := &v1.GetLocationByIPResponse{}

	if net.ParseIP(req.Ip) == nil {
		return resp, status.Error(codes.InvalidArgument, fmt.Sprintf("%s is invalid", req.Ip))
	}

	rec, err := s.db.Get_all(req.Ip)
	if err != nil {
		meta.WithAttribute(ctx, "server.error", err.Error())

		return resp, status.Error(codes.NotFound, fmt.Sprintf("%s was not found", req.Ip))
	}

	country, err := s.query.FindCountryByName(rec.Country_long)
	if err != nil {
		meta.WithAttribute(ctx, "server.error", err.Error())

		return resp, status.Error(codes.NotFound, fmt.Sprintf("%s was not found", req.Ip))
	}

	resp.Location = &v1.Location{
		Country:   country.Codes.Alpha2,
		Continent: continent.Codes[country.Continent],
	}

	return resp, nil
}
