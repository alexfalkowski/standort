// Package gountries adapts github.com/pariz/gountries to standort's country
// provider contract.
//
// It resolves input country codes through the gountries database and returns the
// ISO-3166 alpha-2 country code with the provider's continent name. The domain
// location service performs the final continent-code normalization.
package gountries
