package geoip2

import (
	"context"
	"net"

	"github.com/IncSW/geoip2"
)

// NewProvider for geoip2.
func NewProvider(cfg *Config) (*Provider, error) {
	reader, err := geoip2.NewCountryReaderFromFile(cfg.GetPath())
	if err != nil {
		return nil, err
	}

	return &Provider{reader: reader}, nil
}

// Provider for geoip2.
type Provider struct {
	reader *geoip2.CountryReader
}

// GetByIP a country.
func (p *Provider) GetByIP(_ context.Context, ip string) (string, error) {
	record, err := p.reader.Lookup(net.ParseIP(ip))
	if err != nil {
		return "", err
	}

	return record.Country.Names["en"], nil
}
