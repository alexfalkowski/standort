[![CircleCI](https://circleci.com/gh/alexfalkowski/standort.svg?style=svg)](https://circleci.com/gh/alexfalkowski/standort)
[![Coverage Status](https://coveralls.io/repos/github/alexfalkowski/standort/badge.svg?branch=master)](https://coveralls.io/github/alexfalkowski/standort?branch=master)

# Standort

Standort provides location based information.

## IP Address

The service allows you to get the location by IP address using the awesome [ip2location](https://github.com/ip2location/ip2location-go).

We have checked in the free version of [ip2location.bin](assets/ip2location.bin)

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- [Ruby](.ruby-version)
- Golang

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
