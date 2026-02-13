package cmd

import (
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/strings"
)

// RegisterServer registers the `server` CLI command.
//
// The command starts the standort server using the dependency injection graph
// rooted at `Module`.
//
// The added input is intentionally empty, which delegates configuration loading
// to go-service defaults (for example, loading config via standard flags / env
// supported by the framework).
func RegisterServer(command cli.Commander) {
	cmd := command.AddServer("server", "Start standort server", Module)

	cmd.AddInput(strings.Empty)
}
