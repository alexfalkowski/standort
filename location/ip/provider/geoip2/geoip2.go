package geoip2

import (
	"context"
	"embed"
	"net"
	"strings"

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
	if err != nil {
		return "", p.error(err)
	}

	return record.Country.ISOCode, nil
}

func (p *Provider) error(err error) error {
	if strings.Contains(err.Error(), "invalid") {
		return err
	}

	return errors.ErrNotFound
}
