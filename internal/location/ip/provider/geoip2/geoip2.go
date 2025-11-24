package geoip2

import (
	"embed"
	"net"

	"github.com/IncSW/geoip2"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/runtime"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// NewProvider for geoip2.
func NewProvider(fs embed.FS) *Provider {
	c, err := fs.ReadFile("geoip2.mmdb")
	runtime.Must(err)

	reader, err := geoip2.NewCountryReader(c)
	runtime.Must(err)

	return &Provider{reader: reader}
}

// Provider for geoip2.
type Provider struct {
	reader *geoip2.CountryReader
}

// GetByIP a country.
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
