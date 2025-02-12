package main

import (
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/standort/internal/cmd"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "env:STANDORT_CONFIG_FILE")
	c.AddServer("server", "Start standort server", cmd.ServerOptions...)

	return c
}
