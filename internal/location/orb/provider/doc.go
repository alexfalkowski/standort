// Package provider defines the latitude/longitude lookup contract used by the
// location domain service.
//
// Implementations perform geospatial lookup for coordinates expressed in
// degrees and return provider-level country and continent data. The domain
// service normalizes the returned continent name before exposing it through the
// API.
package provider
