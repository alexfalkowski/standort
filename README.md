[![CircleCI](https://circleci.com/gh/alexfalkowski/standort.svg?style=svg)](https://circleci.com/gh/alexfalkowski/standort)
[![Coverage Status](https://coveralls.io/repos/github/alexfalkowski/standort/badge.svg?branch=master)](https://coveralls.io/github/alexfalkowski/standort?branch=master)

# Standort

Standort provides location based information.

## IP Address

The service allows you to get the location by IP address using the following methods:
- [ip2location](https://github.com/ip2location/ip2location-go)
- [geoip2](https://github.com/IncSW/geoip2)

### IP2Location

This is configured using the following:

```yaml
location:
  ip:
    kind: ip2location
    ip2location:
      path: ../assets/ip2location.bin
```

*NOTE:* This is only needed if you would like to override it. As the system has all this configured.

### GeoIP2

```yaml
location:
  ip:
    kind: geoip2
    geoip2:
      path: ../assets/geoip2.mmdb
```

*NOTE:* This is only needed if you would like to override it. As the system has all this configured.

## Countries/Continents

The service allows you to get the location based on [latitude/longitude](https://www.latlong.net/) using the power of [free maps](https://geojson-maps.ash.ms/).

This is configured using the following:

```yaml
location:
  continent:
    africa_path: ../assets/africa.geojson
    north_america_path: ../assets/north_america.geojson
    oceania_path: ../assets/oceania.geojson
    asia_path: ../assets/asia.geojson
    europe_path: ../assets/europe.geojson
    south_america_path: ../assets/south_america.geojson
```

*NOTE:* This is only needed if you would like to override it. As the system has all this configured.

## Versions

The service contains 2 versions of the API. These are as follows:
- V1 - has separate endpoints for IP and countries.
- V2 - combines both into one endpoint and uses the following headers:
  - For IP address we will use [X-Forwarded-For](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For) header.
  - For geo location we would use [Geolocation](http://tools.ietf.org/html/draft-thomson-geopriv-http-geolocation-00) header. The format is described as the [geo URI](https://www.rfc-editor.org/rfc/rfc5870).

## Client

The client can be used in other projects. This is configured as follows:

```yaml
client:
  v1:
    host: server_host
    timeout: 1s
  v2:
    host: server_host
    timeout: 1s
```

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- [Ruby](.ruby-version)
- Golang

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

The get yourself setup, please run the following:

```sh
make setup
```

### Binaries

To make sure everything compiles for the app, please run the following:

```sh
make build-test
```

### Features

To run all the features, please run the following:

```sh
make features
```

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
