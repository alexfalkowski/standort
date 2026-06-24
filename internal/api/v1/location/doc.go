// Package location adapts domain lookup results into v1 responses.
//
// The package keeps v1 response construction in one place while gRPC and HTTP
// transports remain responsible for their own status mapping.
package location
