package assets

import "github.com/alexfalkowski/go-service/v2/di"

// Module wires embedded lookup asset file metadata.
var Module = di.Module(
	di.Constructor(NewFiles),
)
