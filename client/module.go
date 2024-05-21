package client

import (
	v1 "github.com/alexfalkowski/standort/client/v1"
	v2 "github.com/alexfalkowski/standort/client/v2"
	"go.uber.org/fx"
)

// Module for fx.
var Module = fx.Options(
	v1.Module,
	v2.Module,
)
