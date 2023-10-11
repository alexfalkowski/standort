FROM golang:1.21.3-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -ldflags="-X 'github.com/alexfalkowski/standort/cmd.Version=${version}'" -a -o standort main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app/standort /standort
COPY --from=build /app/assets /assets
ENTRYPOINT ["/standort"]
