package diagnostics

import (
	"maps"

	"github.com/alexfalkowski/go-service/v2/errors"
)

const (
	locationIPError     = "location-ip-error"
	locationLatLngError = "location-lat-lng-error"
	locationPointError  = "location-point-error"
)

// Code is a safe diagnostic code for transport headers or trailers.
type Code string

// NotFound reports a lookup miss.
const NotFound Code = "not_found"

// InvalidPoint reports latitude/longitude outside the supported coordinate range.
const InvalidPoint Code = "invalid_point"

// InvalidGeoURI reports malformed geolocation metadata.
const InvalidGeoURI Code = "invalid_geo_uri"

// IPError adds an IP lookup diagnostic.
func IPError(values Values, code Code) Values {
	return values.with(locationIPError, code)
}

// LatLngError adds a latitude/longitude lookup diagnostic.
func LatLngError(values Values, code Code) Values {
	return values.with(locationLatLngError, code)
}

// PointError adds a geolocation metadata parsing diagnostic.
func PointError(values Values, code Code) Values {
	return values.with(locationPointError, code)
}

// Error wraps err with code-only diagnostic values.
func Error(err error, values Values) error {
	return &diagnosticError{err: err, values: values.copy()}
}

// FromError returns diagnostics carried by err.
func FromError(err error) Values {
	values := Values{}
	values.collect(err)

	return values
}

// Values are safe diagnostics for transport headers or trailers.
type Values map[string]string

// Map returns diagnostics as plain string key/value pairs for transport metadata.
func (v Values) Map() map[string]string {
	values := make(map[string]string, len(v))
	maps.Copy(values, v)

	return values
}

func (v Values) copy() Values {
	return Values(v.Map())
}

func (v Values) with(key string, code Code) Values {
	if v == nil {
		v = Values{}
	}
	v[key] = string(code)

	return v
}

func (v Values) collect(err error) {
	if diagnostic, ok := errors.AsType[*diagnosticError](err); ok {
		maps.Copy(v, diagnostic.values)
	}
}

type diagnosticError struct {
	err    error
	values Values
}

func (d *diagnosticError) Error() string {
	return d.err.Error()
}

func (d *diagnosticError) Unwrap() error {
	return d.err
}
