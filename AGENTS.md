# AGENTS.md

This repository is a Go service called **standort** (location-based information) with two API versions (v1 and v2) exposed over **gRPC** and **HTTP**.

## Recent repo notes (keep in sync)

- README was updated to remove `make setup` (no such target) and to include accurate bootstrap/run steps and gRPC examples.
- Package-level documentation is expected to live in `doc.go` files (a set of `doc.go` files were added across key packages).
- If Go tooling fails with "inconsistent vendoring", run `make dep` to re-vendor (many targets use `-mod vendor`).

## First steps

### Prerequisites (observed)

- Go (see `go.mod` → `go 1.25.0`)
- Git submodules (the repo relies on a `bin/` submodule; see `.gitmodules`)
- Ruby (used for the Ruby test/benchmark harness under `test/`)

Tools referenced by `make` targets (only run if installed): `buf`, `gotestsum`, `govulncheck`, `air`, `hadolint`, `trivy`, `codecovcli`, `mkcert`, `goda`, `dot`, `gsa`, `scc`.

### Bootstrap

```sh
git submodule sync
git submodule update --init
make dep
```

Notes:
- The repo depends on a `bin/` git submodule; initialize it before running most `make` targets.
- Vendoring is used heavily; re-run `make dep` after dependency changes or if you hit vendoring errors.

Notes:
- Most `make` targets call scripts under `./bin/` (submodule). If `bin/` is missing/stale, (re)run the submodule commands above.
- Many Go commands in `make` run with `-mod vendor`; `make dep` runs `go mod vendor`.

## Essential commands

All primary workflow commands are exposed via `make` at repo root.

### Build

```sh
make build
```

- Produces a binary named after the repo directory (see `bin/build/make/grpc.mak:209-211`).

### Run (local)

The `server` command is registered in `main.go` + `internal/cmd/server.go`.

Dev mode (requires `air`):

```sh
make dev
```

Observed from `bin/build/make/grpc.mak:216-218`:
- Runs `air --build.cmd "make dep build"`.
- Runs the binary as `./standort server -i file:test/.config/server.yml`.

Run the built binary directly:

```sh
make build
./standort server -i file:test/.config/server.yml
```

The dev config file `test/.config/server.yml` configures addresses:
- HTTP: `tcp://:11000`
- gRPC: `tcp://:12000`

### Tests

Go unit/spec tests:

```sh
make specs
```

- Uses `gotestsum` directly (see `bin/build/make/grpc.mak:135-137`), and runs Go tests with `-race -mod vendor -coverpkg=...`.

Feature tests (Ruby harness + Go test binary built with `-tags features`):

```sh
make features
```

- `make features` depends on `build-test` (see `bin/build/make/grpc.mak:126-129` and `:212-215`).
- There is a Go test behind the `features` build tag in `main_test.go:1-11`.

Benchmarks (Ruby harness):

```sh
make benchmarks
```

Tip:
- If `make specs` fails because `gotestsum` is missing, install it (it’s invoked directly by the make target).

### Lint / format

```sh
make lint
make fix-lint
make format
```

Observed details:
- Go lint runs field-alignment (`bin/build/go/fa`) + golangci-lint wrapper (`bin/build/go/lint`) (see `bin/build/make/grpc.mak:18-35`).
- golangci-lint is run with `--build-tags features` (see `bin/build/make/grpc.mak:24-28`).
- Ruby lint/format runs inside `test/` via bundler (`make -C test ...`; see `bin/build/make/grpc.mak:36-43` and `test/bin/build/make/ruby.mak`).
- Proto lint/format runs in `api/` via `buf` (`make -C api ...`; see `bin/build/make/grpc.mak:48-55`).

### Protobuf / API generation

```sh
make proto-lint
make proto-format
make proto-generate
```

Observed:
- Protos live under `api/standort/v1` and `api/standort/v2`.
- `api/Makefile` includes `bin/build/make/buf.mak`.
- `buf generate` outputs:
  - Go protobuf + gRPC stubs into this repo (see `api/buf.gen.yaml:3-10`).
  - Ruby protobuf + gRPC stubs into `test/lib` (see `api/buf.gen.yaml:11-14`).

Breaking-change check:

```sh
make proto-breaking
```

- Uses `buf breaking --against 'https://github.com/alexfalkowski/$(NAME).git#branch=master,subdir=api'` (see `api/Makefile` → `bin/build/make/buf.mak:27-29`).

### Security / containers (optional)

```sh
make sec
make trivy-repo
make build-docker platform=amd64
make trivy-image platform=amd64
```

- `make sec` runs `govulncheck -show verbose -test ./...` (see `bin/build/make/grpc.mak:201-207`).

## Repository layout (observed)

- `main.go`: CLI entrypoint; registers the `server` command via `go-service` CLI (`main.go:9-15`, `internal/cmd/server.go:8-13`).
- `internal/cmd/`: DI wiring; composes modules into the server (`internal/cmd/module.go:15-23`).
- `internal/config/`: service config composition; `Config` embeds `go-service/v2/config.Config` (`internal/config/config.go:8-12`).
- `internal/health/`: health registrations/observers (`internal/health/health.go:14-44`).
- `internal/location/`: location domain logic:
  - IP → country code provider (`internal/location/ip/...`).
  - Country/continent mapping provider (`internal/location/country/...`).
  - Lat/lng point-in-polygon provider based on an R-tree (`internal/location/orb/provider/rtree/...`).
- `internal/api/`:
  - `v1/` and `v2/` modules and transports.
  - HTTP transport maps gRPC methods to HTTP routes via `go-service/v2/net/http/rpc`.
  - `internal/api/location/`: transport-facing location logic (parses request metadata like Geolocation).
- `assets/`: embedded data files (GeoJSON + GeoIP DB) via `embed.FS` (`assets/assets.go:7-13`).
- `api/`: protobuf definitions + buf config.
- `test/`: Ruby feature/benchmark harness and runtime config (`test/.config/server.yml`).
- `vendor/`: vendored Go dependencies (used by many `make` recipes).

## Code patterns & conventions (observed)

### Dependency injection

- Modules are composed via `di.Module(...)` (`internal/*/module.go`, `assets/module.go`).
- Constructors are registered using `di.Constructor(...)` (e.g., `assets.NewFS`, `location.New`, `api/location.NewLocator`).

### Transports

- gRPC servers implement generated service interfaces under `internal/api/v*/transport/grpc/`.
- HTTP transport registers routes using generated full method name constants:
  - v1: `internal/api/v1/transport/http/http.go:9-13`
  - v2: `internal/api/v2/transport/http/http.go:9-12`
- Error mapping: the gRPC transport `Server.error` converts **any non-nil error** to `codes.NotFound` (`internal/api/v2/transport/grpc/grpc.go:27-33` and similarly in v1).

### Request metadata

- v2 `Locator` can read inputs from metadata when the request doesn’t provide them:
  - IP: `meta.IPAddr(ctx).Value()` (`internal/api/location/location.go:86-92`).
  - Geolocation header: `meta.Geolocation(ctx)` parsed as a geo URI (`internal/api/location/location.go:99-110`).
- v2 `Locator` records lookup errors as metadata attributes (e.g., `locationIpError`, `locationLatLngError`) and only returns `ErrNotFound` if **both** IP and GEO lookups are missing (`internal/api/location/location.go:53-83`).
- Responses include request metadata via `meta.CamelStrings(ctx, "")` in the gRPC handlers:
  - v1: `internal/api/v1/transport/grpc/location.go:10-30`
  - v2: `internal/api/v2/transport/grpc/location.go:11-21`

### Location lookup implementation

- IP path: GeoIP (`assets/geoip2.mmdb`) → ISO country code → country/continent provider.
- Lat/Lng path: point-in-polygon search using an R-tree populated from `assets/earth.geojson` (`internal/location/orb/provider/rtree/rtree.go:56-72`).

### Formatting

`.editorconfig` enforces:
- 2-space indentation for general files.
- tabs for `Makefile`.
- tabs for `*.go`.

## CI signals (CircleCI)

CircleCI runs (see `.circleci/config.yml`):
- `make dep`, `make lint`, `make proto-breaking`, `make sec`, `make trivy-repo`, `make features`, `make benchmarks`, `make analyse`, `make coverage`, `make codecov-upload`.

If you’re trying to match CI locally, start with:

```sh
make dep
make lint
make specs
```

## Common gotchas (observed)

- **Submodule dependency**: root `Makefile` includes make fragments from `bin/build/make/*` (`Makefile:1-3`). Initialize `bin/` before using most `make` targets.
- **Vendoring**: many targets use `-mod vendor`. After dependency changes, re-run `make dep`.
- **Build tags**: feature harness test is behind `//go:build features` (`main_test.go:1-11`).
- **`make specs` requires `gotestsum`**: it’s invoked directly in the Make target (`bin/build/make/grpc.mak:135-137`).
- **Generated protobuf**: generated `*.pb.go` / `*_grpc.pb.go` are excluded from lint/format (`.golangci.yml:25-43`). Don’t hand-edit generated files; re-run proto generation instead.
- **README mismatch**: historically `README.md` suggested `make setup`, but there is no `setup` target. The README has been updated; if this line reappears, fix the README and/or add the missing target.
- **Package docs**: prefer package-level docs in `doc.go` files. If you add new packages or public surface area, create/update `doc.go` rather than scattering package overview docs across implementation files.
