package continent

// Codes maps continent names to two-letter continent codes.
//
// Keys are the continent names as emitted by the configured country / geo providers
// (for example, "North America").
//
// Values are the corresponding standard two-letter continent codes:
//
//   - AF: Africa
//   - NA: North America
//   - OC: Oceania
//   - AS: Asia
//   - EU: Europe
//   - SA: South America
//
// This map is used by the domain `internal/location` service to normalize provider
// outputs into stable API codes.
var Codes = map[string]string{
	"Africa":        "AF",
	"North America": "NA",
	"Oceania":       "OC",
	"Asia":          "AS",
	"Europe":        "EU",
	"South America": "SA",
}
