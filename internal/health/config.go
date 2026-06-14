package health

import "github.com/alexfalkowski/go-service/v2/time"

// Config configures the health subsystem.
//
// Both fields are required to be positive durations and are parsed from values
// such as "250ms", "5s", or "1m". A zero value does not disable health checks
// and does not act as a default.
//
// The values are used by the health module when registering checks and
// observers:
//
//   - Duration: the interval between health check executions.
//   - Timeout: the maximum amount of time a health check is allowed to run.
type Config struct {
	// Duration is the positive period between health check runs.
	Duration time.Duration `yaml:"duration,omitempty" json:"duration,omitempty" toml:"duration,omitempty" validate:"gt=0"`

	// Timeout is the positive maximum time allowed for a health check execution.
	Timeout time.Duration `yaml:"timeout,omitempty" json:"timeout,omitempty" toml:"timeout,omitempty" validate:"gt=0"`
}
