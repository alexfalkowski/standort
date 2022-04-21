package country

import (
	"github.com/pariz/gountries"
)

// NewQuery for country.
func NewQuery() *gountries.Query {
	return gountries.New()
}
