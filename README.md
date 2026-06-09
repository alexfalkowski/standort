![Gopher](assets/gopher.png)
[![CircleCI](https://circleci.com/gh/alexfalkowski/standort.svg?style=shield)](https://circleci.com/gh/alexfalkowski/standort)
[![codecov](https://codecov.io/gh/alexfalkowski/standort/graph/badge.svg?token=JJP65DPD1M)](https://codecov.io/gh/alexfalkowski/standort)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/standort/v2)](https://goreportcard.com/report/github.com/alexfalkowski/standort/v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/standort/v2.svg)](https://pkg.go.dev/github.com/alexfalkowski/standort/v2)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# 📍 Standort

Standort is a Go service that provides location-based information (country + continent) from:

- an **IP address** (GeoIP2 database lookup), and/or
- a **latitude/longitude point** (point-in-polygon lookup over embedded GeoJSON).

It exposes **two API versions (v1 and v2)** over **gRPC** and **HTTP**. It is useful when an application needs a small service boundary for country/continent enrichment without calling an external provider at request time.

> [!NOTE]
> Standort ships its lookup data as embedded assets, so runtime responses are determined by the GeoIP2 database and GeoJSON files committed under `assets/`.

---

## 🧭 What “location” means

Standort returns:

- `country`: ISO-3166 alpha-2 code (e.g. `US`, `DE`)
- `continent`: two-letter continent code (e.g. `NA`, `EU`)

Lookups are performed using embedded assets:

- `assets/geoip2.mmdb`: IP → country code (GeoIP2)
- `assets/earth.geojson`: lat/lng → country + continent name (GeoJSON polygons), indexed with an R-tree

> [!CAUTION]
> Location accuracy is only as current as the embedded datasets. Update the assets and rebuild the service when lookup freshness matters.

---

## 🔌 API versions

### 1️⃣ v1

v1 has separate RPCs for IP-based lookup and lat/lng-based lookup.

- `GetLocationByIP`
- `GetLocationByLatLng`

### 2️⃣ v2

v2 combines both inputs into a single RPC:

- `GetLocation`

v2 supports passing inputs either directly in the request *or* via request metadata:

- IP address can be derived from request metadata (commonly `X-Forwarded-For`).
- Geolocation can be derived from a `Geolocation` header containing a `geo:` URI (RFC 5870).

If a lookup fails, v2 records error details into response `meta` attributes where possible, and only returns “not found” when neither IP nor GEO yields a location.

> [!WARNING]
> Treat forwarded IP metadata as trusted only after your proxy or gateway has normalized it. Standort reads the metadata supplied by the transport layer; it does not decide whether a forwarded client IP is trustworthy.

---

## 🚀 Quickstart (development)

### ✅ Prerequisites

- Go `1.26.0` (see `go.mod`)
- Git submodules (this repo relies on a `bin/` submodule)
- Ruby (used by the feature/benchmark harness under `test/`)

Optional tools used by specific targets include `air`, `buf`, `gotestsum`, `govulncheck`, `trivy`, and `hadolint`.

### 🧰 Bootstrap

Initialize submodules and vendor dependencies:

```sh
git submodule sync
git submodule update --init
make dep
```

> [!IMPORTANT]
> Most workflow commands include make fragments from the `bin/` submodule, and many Go targets run with `-mod vendor`. If `bin/` is missing or Go reports “inconsistent vendoring,” run the bootstrap commands above again.

### 🏗️ Build

Build the server binary:

```sh
make build
```

### 🖥️ Run locally (dev config)

The repository includes a dev config at `test/.config/server.yml` with default addresses:

- HTTP: `:11000`
- gRPC: `:12000`

#### ♻️ Option A: dev mode (requires `air`)

```sh
make dev
```

This runs `air` to rebuild and restart using the dev config.

#### ▶️ Option B: run the binary directly

If you’ve built the binary (via `make build`), run:

```sh
./standort server -config file:test/.config/server.yml
```

---

## ❤️ Health endpoints

When running with the dev config, the health HTTP observer endpoints are:

- `GET http://localhost:11000/standort/healthz`
- `GET http://localhost:11000/standort/livez`
- `GET http://localhost:11000/standort/readyz`
- `GET http://localhost:11000/standort/metrics`

> [!NOTE]
> go-service prefixes operational HTTP routes with the service name. The dev and test harness service name is `standort`.

The `healthz` endpoint uses go-health's online checker. The `livez` and `readyz` endpoints are local noop observers. The `metrics` endpoint is available with the dev config because it sets `telemetry.metrics.kind: prometheus`.

---

## 📡 API usage examples

The service is gRPC-first; HTTP is wired by routing HTTP requests to the gRPC handlers through go-service RPC routing.

> [!TIP]
> Use `-plaintext` with `grpcurl` for the local dev config unless you have configured TLS.

### 🧪 gRPC examples

You can use `grpcurl` against the dev gRPC address (`localhost:12000`).

#### 🌐 v1: lookup by IP

```sh
grpcurl -plaintext \
  -d '{"ip":"8.8.8.8"}' \
  localhost:12000 \
  standort.v1.Service/GetLocationByIP
```

#### 🗺️ v1: lookup by lat/lng

```sh
grpcurl -plaintext \
  -d '{"lat":52.5200,"lng":13.4050}' \
  localhost:12000 \
  standort.v1.Service/GetLocationByLatLng
```

#### 🌐 v2: lookup with explicit IP

```sh
grpcurl -plaintext \
  -d '{"ip":"8.8.8.8"}' \
  localhost:12000 \
  standort.v2.Service/GetLocation
```

#### 🗺️ v2: lookup with explicit point

```sh
grpcurl -plaintext \
  -d '{"point":{"lat":52.5200,"lng":13.4050}}' \
  localhost:12000 \
  standort.v2.Service/GetLocation
```

#### 🧾 v2: lookup using metadata (headers)

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

Example with a forwarded IP:

```sh
grpcurl -plaintext \
  -H 'X-Forwarded-For: 8.8.8.8' \
  -d '{}' \
  localhost:12000 \
  standort.v2.Service/GetLocation
```

### 🌍 HTTP examples

HTTP runs on `localhost:11000` with the dev config. The RPC router exposes POST routes using the generated gRPC full method names.

#### 🌐 v1: lookup by IP

```sh
curl -sS \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{"ip":"8.8.8.8"}' \
  http://localhost:11000/standort.v1.Service/GetLocationByIP
```

#### 🗺️ v1: lookup by lat/lng

```sh
curl -sS \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{"lat":52.5200,"lng":13.4050}' \
  http://localhost:11000/standort.v1.Service/GetLocationByLatLng
```

#### 🌐 v2: lookup with explicit IP

```sh
curl -sS \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{"ip":"8.8.8.8"}' \
  http://localhost:11000/standort.v2.Service/GetLocation
```

#### 🗺️ v2: lookup with explicit point

```sh
curl -sS \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{"point":{"lat":52.5200,"lng":13.4050}}' \
  http://localhost:11000/standort.v2.Service/GetLocation
```

#### 🧾 v2: lookup using metadata headers

```sh
curl -sS \
  -X POST \
  -H 'Content-Type: application/json' \
  -H 'Geolocation: geo:52.5200,13.4050' \
  -d '{}' \
  http://localhost:11000/standort.v2.Service/GetLocation
```

---

## 🛠️ Development workflow

### 🎨 Format

```sh
make format
```

### 🔎 Lint

```sh
make lint
```

To apply automatic fixes where available:

```sh
make fix-lint
```

### 🧪 Unit/spec tests (Go)

```sh
make specs
```

`make specs` uses `gotestsum` directly. Install it if you don’t already have it.

### 🧩 Feature tests (Ruby harness)

```sh
make features
```

### ⏱️ Benchmarks (Ruby harness)

```sh
make benchmarks
```

### 🧯 Security checks

```sh
make sec
make trivy-repo
```

### ✅ CI parity

CircleCI initializes submodules, vendors dependencies, then runs linting, protobuf breaking checks, security checks, feature tests, benchmarks, analysis, coverage, and Codecov upload. For a focused local pre-push check, start with:

```sh
make dep
make lint
make specs
```

---

## 🧬 Protobuf / API generation

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

> [!CAUTION]
> Do not hand-edit generated protobuf or gRPC files. Change the `.proto` files and run `make proto-generate`.

---

## 🗂️ Repository layout (high level)

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

## 🆘 Support and maintenance

Use the GitHub issue tracker for bug reports, documentation fixes, and feature requests. Maintainers should keep README command examples aligned with `Makefile`, `.circleci/config.yml`, and the generated API definitions under `api/`.

---

## 📝 Changelog

See `CHANGELOG.md`.
