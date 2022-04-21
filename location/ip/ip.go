package ip

import (
	"context"
	"errors"
	"net"

	"github.com/ip2location/ip2location-go/v9"
	"go.uber.org/fx"
)

var (
	// ErrInvalid for ip.
	ErrInvalid = errors.New("invalid ip")
)

// IsValid ip address.
func IsValid(ip string) error {
	if net.ParseIP(ip) == nil {
		return ErrInvalid
	}

	return nil
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
