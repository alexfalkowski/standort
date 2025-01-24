package geoip2

import (
	"context"
	"embed"
	"fmt"
	"net"

	"github.com/IncSW/geoip2"
	"github.com/alexfalkowski/go-service/runtime"
	"github.com/alexfalkowski/standort/location/errors"
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

	return p.code(record), p.error(ip, err)
}

func (p *Provider) error(ip string, err error) error {
	if err != nil {
		return fmt.Errorf("%v: %w", ip, errors.ErrNotFound)
	}

	return nil
}

func (p *Provider) code(record *geoip2.CountryResult) string {
	if record == nil {
		return ""
	}

	return record.Country.ISOCode
}
