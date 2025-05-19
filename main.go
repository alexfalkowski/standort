package main

import (
	"context"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/standort/internal/cmd"
)

var app = sc.NewApplication(func(command *sc.Command) {
	cmd.RegisterServer(command)
})

func main() {
	app.ExitOnError(context.Background())
}
