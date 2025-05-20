package main

import (
	"context"

	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/standort/internal/cmd"
)

var app = cli.NewApplication(func(command cli.Commander) {
	cmd.RegisterServer(command)
})

func main() {
	app.ExitOnError(context.Background())
}
