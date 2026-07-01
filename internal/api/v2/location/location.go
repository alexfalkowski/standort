package location

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/net/grpc/codes"
	"github.com/alexfalkowski/go-service/v2/strings"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/location"
	"google.golang.org/genproto/googleapis/rpc/status"
)

const maxLookups = 100

// ErrTooManyLookups is returned when a batch request contains more than 100
// lookup entries.
var ErrTooManyLookups = errors.New("too many lookups")

// NewLocator constructs a v2 response locator.
func NewLocator(locator *location.Locator) *Locator {
	return &Locator{locator: locator}
}

// Locator resolves v2 location requests and builds generated v2 responses.
type Locator struct {
	locator *location.Locator
}

// Locate resolves a v2 request.
func (l *Locator) Locate(ctx context.Context, req *v2.GetLocationRequest) (*v2.GetLocationResponse, error) {
	locations, err := l.locator.Locate(ctx, req.GetIp(), toPoint(req.GetPoint()))
	if err != nil {
		return nil, err
	}

	return &v2.GetLocationResponse{
		Meta: meta.CamelStrings(ctx, strings.Empty),
		Ip:   toLocation(locations.IP),
		Geo:  toLocation(locations.GEO),
	}, nil
}

// Lookup resolves a v2 batch request.
//
// Responses preserve request order. Each lookup entry is resolved independently;
// entries that resolve at least one location receive a `locations` outcome, and
// entries that do not resolve any location receive a per-entry `NotFound`
// `google.rpc.Status` outcome while the batch response still succeeds. The only
// request-level error currently returned by this method is `ErrTooManyLookups`
// when the request contains more than 100 entries.
func (l *Locator) Lookup(ctx context.Context, req *v2.LookupLocationsRequest) (*v2.LookupLocationsResponse, error) {
	lookups := req.GetLookups()
	if len(lookups) > maxLookups {
		return nil, ErrTooManyLookups
	}

	resp := &v2.LookupLocationsResponse{
		Meta:    meta.CamelStrings(ctx, strings.Empty),
		Lookups: make([]*v2.LocationLookupResponse, 0, len(lookups)),
	}

	for _, lookup := range lookups {
		locations, err := l.locator.Locate(ctx, lookup.GetIp(), toPoint(lookup.GetPoint()))
		if err != nil {
			resp.Lookups = append(resp.Lookups, &v2.LocationLookupResponse{
				Outcome: &v2.LocationLookupResponse_Status{
					Status: &status.Status{
						Code:    int32(codes.NotFound),
						Message: location.ErrNotFound.Error(),
					},
				},
			})
			continue
		}

		resp.Lookups = append(resp.Lookups, &v2.LocationLookupResponse{
			Outcome: &v2.LocationLookupResponse_Locations{
				Locations: &v2.LocationLookupResponse_ResolvedLocations{
					Ip:  toLocation(locations.IP),
					Geo: toLocation(locations.GEO),
				},
			},
		})
	}

	return resp, nil
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
