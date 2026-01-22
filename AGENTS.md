# AGENTS.md

This repository is a Go service called **standort** (location-based information) with two API versions (v1 and v2) exposed over gRPC and HTTP.

## Quick start

### Prerequisites (observed)

- Go (see `go.mod` → `go 1.25.0`)
- Git submodules (the repo relies on a `bin/` submodule; see `.gitmodules`)
- Ruby (used for `make help` rendering and for the Ruby test harness under `test/`)

Other tools are referenced by Make targets (only run them if installed): `buf`, `gotestsum`, `govulncheck`, `air`, `hadolint`, `trivy`, `codecovcli`, `mkcert`, `goda`, `dot`, `gsa`, `scc`.

### Bootstrap

```sh
git submodule sync
git submodule update --init
make dep
```

Notes:
- Most `make` targets rely on scripts in `./bin/` (submodule). If `bin/` is missing or stale, run the submodule commands above.

## Essential commands

All primary workflow commands are exposed via `make` at repo root.

### Build / run

```sh
make build
```

- Produces a binary named after the repo directory (see `bin/build/make/grpc.mak`).

Dev mode (requires `air`):

```sh
make dev
```

- Uses the config file at `test/.config/server.yml` (see `bin/build/make/grpc.mak:216-218`).

### Tests

Go specs (uses `gotestsum` if installed):

```sh
make specs
```

Feature test flow:

```sh
make features
```

- `make features` depends on `build-test` (Go test binary built with `-tags features`; see `bin/build/make/grpc.mak:126-129` and `:212-215`).
- There is a Go test behind the `features` build tag in `main_test.go:1-11`.

### Lint / format

```sh
make lint
make fix-lint
make format
```

Details (observed):
- Go lint uses `bin/build/go/fa` (field alignment) + `bin/build/go/lint` (golangci-lint wrapper) (see `bin/build/make/grpc.mak:18-35`).
- `.golangci.yml` enables a wide set of linters and also enables formatters (gofmt/gofumpt/goimports/gci) while excluding generated protobuf files (`.*\.pb*`).
- Ruby lint/format is executed via `make -C test ...`.
- Proto lint/format is executed via `make -C api ...` and uses `buf`.

### Protobuf / API generation

```sh
make proto-lint
make proto-format
make proto-generate
```

- Protos live under `api/standort/v1` and `api/standort/v2`.
- `buf generate` outputs:
  - Go protobuf + gRPC stubs into the repo (see `api/buf.gen.yaml`).
  - Ruby protobuf + gRPC stubs into `test/lib` (see `api/buf.gen.yaml:11-14`).

Breaking-change check:

```sh
make proto-breaking
```

- Uses `buf breaking --against ...#branch=master,subdir=api` (see `bin/build/make/buf.mak:28-29`).

### Security / containers (optional)

```sh
make sec
make trivy-repo
make build-docker platform=amd64
make trivy-image platform=amd64
```

- `make sec` runs `govulncheck` across the module (see `bin/build/make/go.mak:95-98` and `bin/build/make/grpc.mak:201-207`).

## Repository layout (observed)

- `main.go`: CLI entrypoint; registers the `server` command via `go-service` CLI (see `main.go:9-15`, `internal/cmd/server.go:8-13`).
- `internal/`: application code.
  - `internal/cmd/`: DI wiring (fx-style via `go-service/v2/di`), composes modules (see `internal/cmd/module.go:15-23`).
  - `internal/api/`:
    - `v1/` and `v2/` modules and transports.
    - HTTP transport maps gRPC methods to HTTP routes via `go-service/v2/net/http/rpc` (see `internal/api/v2/transport/http/http.go:9-12`).
  - `internal/location/`: location domain logic (IP lookup + lat/lng polygon search).
  - `internal/config/`: service config composition (`Config` embeds `go-service` config) (see `internal/config/config.go:8-12`).
  - `internal/health/`: health endpoints and registrations (see `internal/health/health.go`).
- `assets/`: embedded data files (GeoJSON + GeoIP DB) via `embed.FS` (see `assets/assets.go:7-13`).
- `api/`: protobuf definitions + buf config.
- `test/`: Ruby-based feature/benchmark harness (bundler + nonnative) and runtime config (`test/.config/server.yml`).
- `bin/`: git submodule providing shared build scripts and make includes.

## Code patterns & conventions

### Dependency injection

- Modules are composed via `di.Module(...)` (see matches in `internal/*/module.go`).
- Constructors are registered using `di.Constructor(...)`; hooks/registrations via `di.Register(...)`.

### Transports

- gRPC servers implement generated service interfaces and convert domain errors to gRPC status (currently maps non-nil errors to `codes.NotFound`) (see `internal/api/v2/transport/grpc/grpc.go:27-33`).
- HTTP transport registers routes using the generated full method name constants and points them at the gRPC handler functions (see `internal/api/v1/transport/http/http.go:9-13`).

### Location logic

- IP path: IP → ISO country code via GeoIP (`assets/geoip2.mmdb`) → country/continent via `gountries` (see `internal/location/location.go:26-39`).
- Lat/Lng path: point-in-polygon search using an R-tree populated from `assets/earth.geojson` (see `internal/location/orb/provider/rtree/rtree.go:56-72`).

### Formatting

- `.editorconfig` enforces:
  - 2-space indentation for general files
  - tabs for `Makefile`
  - tabs for `*.go`

## CI signals (CircleCI)

CircleCI runs (see `.circleci/config.yml`):
- `make dep`, `make lint`, `make proto-breaking`, `make sec`, `make trivy-repo`, `make features`, `make benchmarks`, `make analyse`, `make coverage`.

If you’re trying to match CI locally, start with:

```sh
make dep
make lint
make specs
```

## Common gotchas

- **Submodule dependency**: root `Makefile` includes make fragments from `bin/build/make/*` (submodule). Always ensure `bin/` is initialized.
- **Vendoring**: the build/test recipes use `-mod vendor` and `make dep` runs `go mod vendor`; if dependencies look “missing”, re-run `make dep`.
- **Build tags**: some tests (e.g. `main_test.go`) are behind `//go:build features`; they won’t run under plain `go test ./...` without `-tags features`.
- **Ruby lint uses bundler + native gems**: `make ruby-lint` runs `bundler exec rubocop` inside `test/` (see `bin/build/make/ruby.mak:3-6`). If you hit a Ruby `LoadError` / missing `libruby*.dylib`, re-run `make ruby-dep` (or `make -C test dep`) to rebuild `test/vendor/bundle` for your local Ruby.
- **Generated protobuf**: protobuf outputs (`*.pb.go`, `*_grpc.pb.go`) are treated as generated and are excluded from lint/format in `.golangci.yml`.
