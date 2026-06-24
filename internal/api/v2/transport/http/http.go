package http

import (
	"github.com/alexfalkowski/go-service/v2/context"
	httpmeta "github.com/alexfalkowski/go-service/v2/net/http/meta"
	"github.com/alexfalkowski/go-service/v2/net/http/rpc"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/api/v2/location"
	"github.com/alexfalkowski/standort/v2/internal/diagnostics"
)

// Register registers the v2 HTTP routes for the Standort API.
//
// This package uses go-service's RPC routing, mapping the generated full method
// name to the corresponding HTTP handler:
//
//   - `standort.v2.Service/GetLocation` → `getLocation`
//
// The HTTP server and route shapes are provided by `rpc.Route`; this function only
// wires the route to the v2 locator.
func Register(locator *location.Locator) {
	rpc.Route(v2.Service_GetLocation_FullMethodName, getLocation(locator))
}

func setFailureHeaders(ctx context.Context, values diagnostics.Values) {
	headers := httpmeta.Response(ctx).Header()
	for key, value := range values {
		headers.Set(key, value)
	}
}
