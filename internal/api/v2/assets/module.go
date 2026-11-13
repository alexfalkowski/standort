package assets

import (
	"github.com/alexfalkowski/go-signal/v2"
	"github.com/alexfalkowski/standort/v3/internal/assets"
)

// Module wires the v2 embedded lookup asset repository.
var Module = signal.Module(
	assets.Module,
	signal.Constructor(NewRepository),
)
