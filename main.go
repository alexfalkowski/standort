package main

import (
	"os"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/standort/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput("env:STANDORT_CONFIG_FILE")
	c.AddServer(cmd.ServerOptions...)

	return c
}
