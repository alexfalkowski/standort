package location

import (
	"github.com/alexfalkowski/standort/location/continent"
	"github.com/alexfalkowski/standort/location/ip"
)

// Config for location.
type Config struct {
	Continent continent.Config `yaml:"continent" json:"continent"`
	IP        ip.Config        `yaml:"ip" json:"ip"`
}
