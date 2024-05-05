package ip2location

import (
	"bytes"
	"context"
	"embed"

	"github.com/alexfalkowski/go-service/runtime"
	"github.com/ip2location/ip2location-go/v9"
	"go.uber.org/fx"
)

// NewProvider for ip2location.
func NewProvider(lc fx.Lifecycle, fs embed.FS) *Provider {
	c, err := fs.ReadFile("ip2location.bin")
	runtime.Must(err)

	b := &buffer{Reader: bytes.NewReader(c)}

	db, err := ip2location.OpenDBWithReader(b)
	runtime.Must(err)

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			db.Close()

			return nil
		},
	})

	return &Provider{db: db}
}

// Provider for ip2location.
type Provider struct {
	db *ip2location.DB
}

// GetByIP a country.
func (p *Provider) GetByIP(_ context.Context, ip string) (string, error) {
	rec, err := p.db.Get_all(ip)
	if err != nil {
		return "", err
	}

	return rec.Country_short, nil
}

type buffer struct {
	*bytes.Reader
}

func (b *buffer) Close() error {
	return nil
}
