package assets

import "github.com/alexfalkowski/go-signal/v2"

// Module wires embedded lookup asset file metadata.
var Module = signal.Module(
	signal.Constructor(NewFiles),
)
