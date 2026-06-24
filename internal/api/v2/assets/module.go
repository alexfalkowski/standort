package assets

import (
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/standort/v2/internal/assets"
)

// Module wires the v2 embedded lookup asset repository.
var Module = di.Module(
	assets.Module,
	di.Constructor(NewRepository),
)
