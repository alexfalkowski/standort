package main

import (
	"os"
	"time"

	scmd "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/standort/cmd"
)

// nolint:gomnd
func main() {
	command := scmd.New(15 * time.Second)

	command.AddServer(cmd.ServerOptions)
	command.AddWorker(cmd.WorkerOptions)
	command.AddVersion(cmd.Version)

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
