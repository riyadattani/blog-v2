package inmem

import (
	"context"
	"errors"
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
	// pp.PP(">>>>> repository fs", r.filesystem)
	// pp.PP(fmt.Sprintf(">>>>> repository title %q", title))
	entries, err := fs.ReadDir(r.filesystem, ".")
	if err != nil {
		return nil, false, fmt.Errorf("cannot read directory: %v", err)
	}

	// pp.PP(">>>>> repository entries", entries)

	if len(entries) == 0 {
		return nil, false, ErrNoEntriesFound
	}

	// todo: use fs.WalkDir(fsys, ".", myFunc)

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
