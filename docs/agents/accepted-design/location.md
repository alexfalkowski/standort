# Accepted design for location lookups

Part of the [accepted-design index](../accepted-design.md).

## v2 metadata fallback

`GetLocation` and each `LookupLocations` entry resolve IP and geolocation
independently. An explicit IP or point takes precedence for its own lookup.
When the IP is omitted, the locator uses transport IP metadata. When the point
is omitted, the locator uses the `Geolocation` metadata. The separate `ip` and
`geo` response fields identify the results, so an explicit input can produce a
result alongside a result derived from metadata.

Treat this fallback as part of the v2 contract in feature, reliability, and
API design reviews. Revisit an explicit-input-only switch if evidence shows
the contract has changed.
[Shared locator](../../../internal/api/location/location.go) · [v2 response assembly](../../../internal/api/v2/location/location.go)

## Embedded lookup asset freshness

Standort embeds the GeoIP database and GeoJSON lookup data at build time.
Refresh the assets when lookup freshness is needed and rebuild the service.
Revisit adding an automatic freshness guard when there is a concrete stale-data
failure or a documented freshness requirement.
[Asset refresh instructions](../../../README.md) · [Dataset sources](../../../assets/SOURCES.md)
