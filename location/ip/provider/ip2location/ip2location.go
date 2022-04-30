package ip2location

import (
	"context"

	"github.com/ip2location/ip2location-go/v9"
	"go.uber.org/fx"
)

// NewProvider for ip2location.
func NewProvider(lc fx.Lifecycle, cfg *Config) (*Provider, error) {
	db, err := ip2location.OpenDB(cfg.GetPath())
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			db.Close()

			return nil
		},
	})

	return &Provider{db: db}, nil
}

// Provider for ip2location.
type Provider struct {
	db *ip2location.DB
}

// GetByIP a country.
func (p *Provider) GetByIP(ctx context.Context, ip string) (string, error) {
	rec, err := p.db.Get_all(ip)
	if err != nil {
		return "", err
	}

	return rec.Country_long, nil
}
