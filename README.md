![Gopher](assets/gopher.png)
[![CircleCI](https://circleci.com/gh/alexfalkowski/standort.svg?style=shield)](https://circleci.com/gh/alexfalkowski/standort)
[![codecov](https://codecov.io/gh/alexfalkowski/standort/graph/badge.svg?token=JJP65DPD1M)](https://codecov.io/gh/alexfalkowski/standort)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/standort/v2)](https://goreportcard.com/report/github.com/alexfalkowski/standort/v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/standort/v2.svg)](https://pkg.go.dev/github.com/alexfalkowski/standort/v2)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Standort

Standort is a Go service that provides location-based information (country + continent) from:

- an **IP address** (GeoIP2 database lookup), and/or
- a **latitude/longitude point** (point-in-polygon lookup over embedded GeoJSON).

It exposes **two API versions (v1 and v2)** over **gRPC** and **HTTP**.

---

## What “location” means

Standort returns:

- `country`: ISO-3166 alpha-2 code (e.g. `US`, `DE`)
- `continent`: two-letter continent code (e.g. `NA`, `EU`)

Lookups are performed using embedded assets:

- `assets/geoip2.mmdb`: IP → country code (GeoIP2)
- `assets/earth.geojson`: lat/lng → country + continent name (GeoJSON polygons), indexed with an R-tree

---

## API versions

### v1

v1 has separate RPCs for IP-based lookup and lat/lng-based lookup.

- `GetLocationByIP`
- `GetLocationByLatLng`

### v2

v2 combines both inputs into a single RPC:

- `GetLocation`

v2 supports passing inputs either directly in the request *or* via request metadata:

- IP address can be derived from request metadata (commonly `X-Forwarded-For`).
- Geolocation can be derived from a `Geolocation` header containing a `geo:` URI (RFC 5870).

If a lookup fails, v2 records error details into response `meta` attributes where possible, and only returns “not found” when neither IP nor GEO yields a location.

---

## Quickstart (development)

### Prerequisites

- Go (see `go.mod` for the required version)
- Git submodules (this repo relies on a `bin/` submodule)
- Ruby (used by the feature/benchmark harness under `test/`)

### Bootstrap

Initialize submodules and vendor dependencies:

```sh
git submodule sync
git submodule update --init
make dep
```

Notes:
- Many `make` targets run with `-mod vendor`.
- If you see “inconsistent vendoring” errors, re-run `make dep`.

### Build

Build the server binary:

```sh
make build
```

### Run locally (dev config)

The repository includes a dev config at `test/.config/server.yml` with default addresses:

- HTTP: `:11000`
- gRPC: `:12000`

#### Option A: dev mode (requires `air`)

```sh
make dev
```

This runs `air` to rebuild and restart using the dev config.

#### Option B: run the binary directly

If you’ve built the binary (via `make build`), run:

```sh
./standort server -i file:test/.config/server.yml
```

---

## Health endpoints

When running with the dev config, the health HTTP observer endpoints are:

- `GET http://localhost:11000/healthz`
- `GET http://localhost:11000/livez`
- `GET http://localhost:11000/readyz`

(Exact behavior depends on the go-service/go-health wiring.)

---

## API usage examples

The service is gRPC-first; HTTP is wired by routing HTTP requests to the gRPC handlers. Exact HTTP routes and encoding are provided by the go-service RPC router.

### gRPC examples

You can use `grpcurl` against the dev gRPC address (`localhost:12000`).

> Tip: you may need `-plaintext` for local development unless you’ve configured TLS.

#### v1: lookup by IP

```sh
grpcurl -plaintext \
  -d '{"ip":"8.8.8.8"}' \
  localhost:12000 \
  standort.v1.Service/GetLocationByIP
```

#### v1: lookup by lat/lng

```sh
grpcurl -plaintext \
  -d '{"lat":52.5200,"lng":13.4050}' \
  localhost:12000 \
  standort.v1.Service/GetLocationByLatLng
```

#### v2: lookup with explicit IP

```sh
grpcurl -plaintext \
  -d '{"ip":"8.8.8.8"}' \
  localhost:12000 \
  standort.v2.Service/GetLocation
```

#### v2: lookup with explicit point

```sh
grpcurl -plaintext \
  -d '{"point":{"lat":52.5200,"lng":13.4050}}' \
  localhost:12000 \
  standort.v2.Service/GetLocation
```

#### v2: lookup using metadata (headers)

v2 can fall back to metadata when request fields are omitted:

- IP can come from forwarded IP metadata (commonly derived from `X-Forwarded-For`).
- Point can come from a `Geolocation` header using a `geo:` URI.

Example with a geo URI:

```sh
grpcurl -plaintext \
  -H 'Geolocation: geo:52.5200,13.4050' \
  -d '{}' \
  localhost:12000 \
  standort.v2.Service/GetLocation
```

Example with a forwarded IP (exact metadata key handling is framework-dependent, but commonly derived from HTTP `X-Forwarded-For` in the gateway/proxy layer):

```sh
grpcurl -plaintext \
  -H 'X-Forwarded-For: 8.8.8.8' \
  -d '{}' \
  localhost:12000 \
  standort.v2.Service/GetLocation
```

### HTTP examples

HTTP runs on `localhost:11000` with the dev config.

Because HTTP routing is generated/wired via go-service RPC routing, the easiest accurate source of truth for HTTP paths is the service at runtime (logs) or the framework documentation/config. The mapping is based on the generated gRPC full method names:

- v1:
  - `standort.v1.Service/GetLocationByIP`
  - `standort.v1.Service/GetLocationByLatLng`
- v2:
  - `standort.v2.Service/GetLocation`

If you want fully concrete `curl` examples for the HTTP endpoints (paths, methods, and JSON shapes), run the server and inspect the registered routes (or share the route output/logs), and I’ll add them verbatim.

---

## Development workflow

### Format

```sh
make format
```

### Lint

```sh
make lint
```

To apply automatic fixes where available:

```sh
make fix-lint
```

### Unit/spec tests (Go)

```sh
make specs
```

Note: `make specs` uses `gotestsum` directly. Install it if you don’t already have it.

### Feature tests (Ruby harness)

```sh
make features
```

### Benchmarks (Ruby harness)

```sh
make benchmarks
```

---

## Protobuf / API generation

Protos live under `api/standort/v1` and `api/standort/v2`.

To lint and generate code:

```sh
make proto-lint
make proto-format
make proto-generate
```

Breaking-change check:

```sh
make proto-breaking
```

---

## Repository layout (high level)

- `main.go`: CLI entrypoint; registers the `server` command.
- `internal/cmd`: DI composition and command wiring.
- `internal/config`: service config composition.
- `internal/health`: health registrations/observers.
- `internal/location`: domain lookup logic (IP + point-in-polygon).
- `internal/api/v1`, `internal/api/v2`: API modules and transports (gRPC + HTTP wiring).
- `assets`: embedded runtime datasets (GeoJSON + GeoIP DB).
- `api`: protobuf definitions + buf config.
- `test`: Ruby feature/benchmark harness + example config.
- `vendor`: vendored Go dependencies.

---

## Changelog

See `CHANGELOG.md`.