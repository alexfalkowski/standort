package geoip2

import (
	"embed"
	"net"

	"github.com/IncSW/geoip2"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/runtime"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// NewProvider constructs an IP→country provider backed by a GeoIP2 country database.
//
// The provider expects an embedded file named `geoip2.mmdb` to be available in `fs`.
// It reads the database into memory during construction and will terminate the
// process (via `runtime.Must`) if the asset cannot be loaded or the database
// cannot be initialized.
//
// The returned provider resolves IP addresses into ISO-3166 alpha-2 country codes
// (for example "US").
func NewProvider(fs embed.FS) *Provider {
	c, err := fs.ReadFile("geoip2.mmdb")
	runtime.Must(err)

	reader, err := geoip2.NewCountryReader(c)
	runtime.Must(err)

	return &Provider{reader: reader}
}

// Provider resolves an IP address into an ISO-3166 alpha-2 country code using GeoIP2.
type Provider struct {
	reader *geoip2.CountryReader
}

// GetByIP resolves the ISO-3166 alpha-2 country code for the given IP address.
//
// The `ip` parameter must be a textual IPv4 or IPv6 address. The value is parsed
// using `net.ParseIP`. If parsing fails, the lookup is performed with a nil IP,
// and the underlying reader may return an error.
//
// If the database contains no matching record, the returned country code is an
// empty string.
//
// Returns:
//   - countryCode: ISO-3166 alpha-2 code (e.g. "US"), or empty string if not found
//   - err: any error returned by the GeoIP2 reader
func (p *Provider) GetByIP(_ context.Context, ip string) (string, error) {
	record, err := p.reader.Lookup(net.ParseIP(ip))

	return p.code(record), err
}

func (p *Provider) code(record *geoip2.CountryResult) string {
	if record == nil {
		return strings.Empty
	}

	return record.Country.ISOCode
}
