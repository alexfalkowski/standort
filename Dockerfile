FROM golang:1.18.1-bullseye AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o standort main.go

FROM debian:bullseye-slim

WORKDIR /

RUN mkdir -p assets && \
    apt-get update && \
    apt-get -y upgrade && \
    rm -rf /var/lib/apt/lists/*

COPY --from=build /app/standort /standort
COPY --from=build /app/assets/ip2location.bin /assets/ip2location.bin
ENTRYPOINT ["/standort"]
