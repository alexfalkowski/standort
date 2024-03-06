package client

import (
	v1 "github.com/alexfalkowski/standort/client/v1"
	"go.uber.org/fx"
)

// ClientModule for fx.
var ClientModule = fx.Options(
	v1.Module,
)
