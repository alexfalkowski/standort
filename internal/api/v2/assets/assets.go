package assets

import (
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/meta"
	"github.com/alexfalkowski/go-service/v2/strings"
	v2 "github.com/alexfalkowski/standort/v2/api/standort/v2"
	"github.com/alexfalkowski/standort/v2/internal/assets"
)

// NewRepository constructs a v2 embedded lookup asset repository.
func NewRepository(files assets.Files) *Repository {
	return &Repository{files: files}
}

// Repository returns metadata for embedded lookup assets.
type Repository struct {
	files assets.Files
}

// Get returns embedded lookup asset metadata.
func (r *Repository) Get(ctx context.Context) *v2.GetLookupAssetsResponse {
	files := make([]*v2.LookupAsset, 0, len(r.files))
	for _, file := range r.files {
		files = append(files, &v2.LookupAsset{
			Name:              file.Name,
			SizeBytes:         file.SizeBytes,
			ChecksumAlgorithm: file.ChecksumAlgorithm,
			Checksum:          file.Checksum,
		})
	}

	return &v2.GetLookupAssetsResponse{
		Meta:   meta.CamelStrings(ctx, strings.Empty),
		Assets: files,
	}
}
