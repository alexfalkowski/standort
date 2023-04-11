.PHONY: vendor

include bin/build/make/service.mak

# Build release binary.
build:
	go build -race -ldflags="-X 'github.com/alexfalkowski/standort/cmd.Version=latest'" -mod vendor -o standort main.go

# Build test binary.
build-test:
	go test -race -ldflags="-X 'github.com/alexfalkowski/standort/cmd.Version=latest'" -mod vendor -c -tags features -covermode=atomic -o standort -coverpkg=./... github.com/alexfalkowski/standort

# Release to docker hub.
docker:
	bin/build/docker/push standort
