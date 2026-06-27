package assets

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"

	"github.com/alexfalkowski/go-service/v2/runtime"
)

const checksumAlgorithm = "sha256"

var names = []string{"geoip2.mmdb", "earth.geojson"}

// File describes one embedded lookup asset file.
type File struct {
	Name              string
	ChecksumAlgorithm string
	Checksum          string
	SizeBytes         uint64
}

// Files contains metadata for embedded lookup asset files.
type Files []File

// NewFiles constructs metadata for the expected embedded lookup asset files.
//
// It reads `geoip2.mmdb` and `earth.geojson`, computes SHA-256 checksums, and
// records each file's byte size. Construction fails through `runtime.Must` if
// either expected file cannot be read from the embedded filesystem.
func NewFiles(fs embed.FS) Files {
	files := make(Files, 0, len(names))
	for _, name := range names {
		files = append(files, newFile(fs, name))
	}

	return files
}

func newFile(fs embed.FS, name string) File {
	data, err := fs.ReadFile(name)
	runtime.Must(err)

	sum := sha256.Sum256(data)

	return File{
		Name:              name,
		ChecksumAlgorithm: checksumAlgorithm,
		Checksum:          hex.EncodeToString(sum[:]),
		SizeBytes:         uint64(len(data)),
	}
}
