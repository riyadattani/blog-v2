package inmem

import (
	"context"
	"errors"
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

	if len(entries) == 0 {
		return nil, false, ErrNoEntriesFound
	}

	for _, entry := range entries {
		if entry.Name() == title {
			fileBytes, err := fs.ReadFile(r.filesystem, entry.Name())
			if err != nil {
				return nil, found, ErrEmptyEntry
			}

			return fileBytes, true, nil
		}
	}

	return nil, false, ErrEntryNotFound
}

var (
	ErrEntryNotFound  = errors.New("could not find entry in the filesystem")
	ErrNoEntriesFound = errors.New("filesystem is empty")
	ErrEmptyEntry     = errors.New("no data in entry")
)
