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

When replacing embedded lookup data, keep the filenames and formats unchanged:

- `geoip2.mmdb` must remain a GeoIP2 country database compatible with the `github.com/IncSW/geoip2` reader.
- `earth.geojson` must remain a GeoJSON feature collection whose indexed features have Polygon or MultiPolygon geometry, a two-character `iso_a2` or `iso_a2_eh` country code, and a supported `continent` value.

Record the dataset source/provenance in the asset update change, then rebuild and run the feature harness so changed lookup results are visible in review.

---

## 🔌 API versions

### 1️⃣ v1

v1 has separate RPCs for IP-based lookup and lat/lng-based lookup.

- `GetLocationByIP`
- `GetLocationByLatLng`

### 2️⃣ v2

v2 combines both inputs into a single lookup RPC and also supports batch
lookups:

- `GetLocation`
- `LookupLocations`
- `GetLookupAssets`

v2 supports passing inputs either directly in the request *or* via request metadata:

- IP address can be derived from request metadata (commonly `X-Forwarded-For`).
- Geolocation can be derived from a `Geolocation` header containing a `geo:` URI (RFC 5870).

If a lookup fails, v2 keeps trying any other available input, and only returns “not found” when neither IP nor GEO yields a location.

v2 response fields are independent:

- `ip` is populated when the IP lookup succeeds.
- `geo` is populated when the point lookup succeeds.
- both fields are populated when both inputs succeed.
- on partial success, the successful field is returned without failed-side diagnostics.

Terminal lookup failures return gRPC `NotFound`; the HTTP RPC router exposes the same lookup miss as HTTP `404`. v2 transports may attach code-only diagnostics to the terminal error metadata, using `location-ip-error`, `location-lat-lng-error`, or `location-point-error`.

`LookupLocations` accepts up to 100 lookup entries and preserves request order
in the response. Entries that resolve successfully populate `ip`, `geo`, or
both. Entries that do not resolve any location populate a per-entry
`google.rpc.Status` instead of failing the whole batch. Requests with more than
100 lookup entries fail the whole call with gRPC `InvalidArgument`; the HTTP RPC
mapping returns `400`.

`GetLookupAssets` returns read-only metadata for the embedded lookup assets,
including asset name, size in bytes, checksum algorithm, and checksum. It is
informational only; it does not enforce asset freshness.

> [!WARNING]
> Treat forwarded IP metadata as trusted only after your proxy or gateway has normalized it. Standort reads the metadata supplied by the transport layer; it does not decide whether a forwarded client IP is trustworthy.

---

## 🚀 Quickstart (development)

### ✅ Prerequisites

- Go (see `go.mod`)
- Git submodules with GitHub SSH access (this repo relies on a `bin/` submodule)
- Ruby with Bundler (used by the feature/benchmark harness under `test/`)

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
> The `bin` submodule intentionally uses an SSH URL (`git@github.com:alexfalkowski/bin.git`), so fresh environments need GitHub SSH credentials before `git submodule update --init` can succeed.

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

## 🔧 Configuration

`standort server` reads a go-service configuration file from `-config` (or `-c`). The dev sample at `test/.config/server.yml` is the authoritative local example.

Standort-owned configuration currently adds the required `health` section:

```yaml
health:
  duration: 1s
  timeout: 1s
```

Both values must be positive durations. The rest of the file is the embedded go-service configuration, inlined at the top level, including transport addresses/timeouts, limiter settings, logger, metrics, tracer, and ID settings.

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

Use `livez` or `readyz` for local process probes that must not depend on public egress. The `healthz` endpoint is an online check and is expected to report unhealthy when the process cannot reach the public internet.

The gRPC health service registers these service names:

- `standort.v1.Service`
- `standort.v2.Service`

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

#### 📦 v2: batch lookup

```sh
grpcurl -plaintext \
  -d '{"lookups":[{"ip":"8.8.8.8"},{"point":{"lat":52.5200,"lng":13.4050}},{"ip":"192.0.2.1"}]}' \
  localhost:12000 \
  standort.v2.Service/LookupLocations
```

#### 🧾 v2: embedded lookup assets

```sh
grpcurl -plaintext \
  -d '{}' \
  localhost:12000 \
  standort.v2.Service/GetLookupAssets
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

Example with a forwarded IP:

```sh
curl -sS \
  -X POST \
  -H 'Content-Type: application/json' \
  -H 'X-Forwarded-For: 8.8.8.8' \
  -d '{}' \
  http://localhost:11000/standort.v2.Service/GetLocation
```

#### 📦 v2: batch lookup

```sh
curl -sS \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{"lookups":[{"ip":"8.8.8.8"},{"point":{"lat":52.5200,"lng":13.4050}},{"ip":"192.0.2.1"}]}' \
  http://localhost:11000/standort.v2.Service/LookupLocations
```

#### 🧾 v2: embedded lookup assets

```sh
curl -sS \
  -X POST \
  -H 'Content-Type: application/json' \
  -d '{}' \
  http://localhost:11000/standort.v2.Service/GetLookupAssets
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

The Ruby harness is configured by `test/nonnative.yml` and writes logs, coverage, and report artifacts under `test/reports/`.

The harness starts and stops `../standort server` itself with `-config file:.config/server.yml` relative to `test/`, so leave ports `11000` and `12000` free before running it. To rerun a focused slice, pass paths relative to `test/`:

```sh
make features feature=features/v2/transport/grpc/api.feature tags=@grpc
```

### ⏱️ Benchmarks (Ruby harness)

```sh
make benchmarks
```

Benchmark runs use the same harness configuration and report directory as feature tests.

Focused benchmark runs also use paths relative to `test/`:

```sh
make benchmarks feature=features/v2/transport/http/benchmark.feature
```

### 🧯 Security checks

```sh
make sec
make trivy-repo
```

`make sec` is the full local security wrapper: it runs `govulncheck -show verbose -test ./...` and the Trivy repository scan. Use `make trivy-repo` only when you want the narrower Trivy-only scan.

### ✅ CI parity

CircleCI initializes submodules, vendors dependencies, then runs linting,
protobuf breaking and generated-output stale checks, security checks, feature
tests, benchmarks, analysis, coverage, and Codecov upload. To mirror the main
validation sequence locally, use:

```sh
make dep
make lint
make proto-breaking
make proto-stale
make sec
make features
make benchmarks
make analyse
make coverage
```

Use `make specs` as a focused Go test check; it is not a separate CircleCI
step in the main build job.

---

## 🧬 Protobuf / API generation

Protos live under `api/standort/v1` and `api/standort/v2`.

To lint and generate code:

```sh
make proto-lint
make proto-format
make proto-generate
```

`buf generate` writes Go protobuf/gRPC files into this repository and Ruby protobuf/gRPC files under `test/lib`. It uses Buf remote plugins, so `make proto-generate` and `make proto-stale` may need network access on first run or when the plugin cache is empty. When a `.proto` file is renamed or deleted, remove any orphaned generated Go and Ruby outputs in the same change.

Breaking-change check:

```sh
make proto-breaking
```

`make proto-breaking` compares `api/` with the `master` branch on GitHub and requires network access to that remote baseline.

Generated-output freshness check:

```sh
make proto-stale
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
