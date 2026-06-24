// Package diagnostics carries safe lookup failure diagnostics to transports.
//
// Diagnostics are code-only key/value pairs intended for response headers or
// gRPC trailers. They deliberately avoid raw error messages so response bodies
// and transport metadata do not expose provider-specific failure text.
package diagnostics
