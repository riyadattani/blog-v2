package repository

import (
	"context"
	"io/fs"
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Driver interface {
	Get(ctx context.Context, title string) (stuff []byte, found bool, err error)
}

type Specification struct {
	NewRepo     Driver
	MakeContext func(tb testing.TB) context.Context
	MakeDir     func(tb testing.TB) fs.FS
}

func (s Specification) Test(t *testing.T) {
	t.Run("get from dir", func(t *testing.T) {
		ctx := s.MakeContext(t)
		repo := s.NewRepo

		dir := s.MakeDir(t)

		entries, err := fs.ReadDir(dir, ".")
		assert.NoError(t, err)

		for _, entry := range entries {
			_, found, err := repo.Get(ctx, entry.Name())

			assert.NoError(t, err)
			assert.True(t, found)
			// assert.Equal(t, expected, got)
		}
	})
}
