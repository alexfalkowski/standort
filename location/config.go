package location

import (
	"github.com/alexfalkowski/standort/location/ip"
)

// Config for location.
type Config struct {
	IP ip.Config `yaml:"ip"`
}
