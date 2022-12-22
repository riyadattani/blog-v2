package inmem

import (
	"context"
	"fmt"
	"io/fs"
)

type Repository struct {
	filesystem fs.FS
}

func NewRepository(system fs.FS) *Repository {
	return &Repository{
		filesystem: system,
	}
}

func (r *Repository) Get(ctx context.Context, title string) (stuff []byte, found bool, err error) {
	entries, err := fs.ReadDir(r.filesystem, ".")
	if err != nil {
		return nil, false, err
	}

	for _, entry := range entries {
		if entry.Name() == title {
			fileBytes, err := fs.ReadFile(r.filesystem, entry.Name())
			if err != nil {
				return nil, found, fmt.Errorf("found file but there is nothing inside")
			}

			return fileBytes, true, nil
		}
	}

	return nil, false, fmt.Errorf("something went wrong when getting file from repo")
}
