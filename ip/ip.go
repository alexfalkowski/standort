package ip

import (
	"github.com/ip2location/ip2location-go/v9"
)

// NewDB for ip.
func NewDB(cfg *Config) (*ip2location.DB, error) {
	return ip2location.OpenDB(cfg.Path)
}
