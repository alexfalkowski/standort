package ip

import (
	"context"
	"net"

	"github.com/ip2location/ip2location-go/v9"
	"go.uber.org/fx"
)

// IsValid ip address.
func IsValid(ip string) bool {
	return net.ParseIP(ip) != nil
}

// NewDB for ip.
func NewDB(lc fx.Lifecycle, cfg *Config) (*ip2location.DB, error) {
	db, err := ip2location.OpenDB(cfg.Path)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			db.Close()

			return nil
		},
	})

	return db, nil
}
