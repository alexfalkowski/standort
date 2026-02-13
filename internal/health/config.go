package health

// Config configures the health subsystem.
//
// Both fields are parsed as durations (for example: "250ms", "5s", "1m").
//
// The values are used by the health module when registering checks and
// observers:
//
//   - Duration: the interval between health check executions.
//   - Timeout: the maximum amount of time a health check is allowed to run.
type Config struct {
	// Duration is the period between health check runs.
	Duration string `yaml:"duration,omitempty" json:"duration,omitempty" toml:"duration,omitempty"`

	// Timeout is the maximum time allowed for a health check execution.
	Timeout string `yaml:"timeout,omitempty" json:"timeout,omitempty" toml:"timeout,omitempty"`
}
