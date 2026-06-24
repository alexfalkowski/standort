// Package location adapts transport-facing lookup results into v2 responses.
//
// The package keeps v2 response construction in one place while gRPC and HTTP
// transports remain responsible for their own status and diagnostic metadata.
package location
