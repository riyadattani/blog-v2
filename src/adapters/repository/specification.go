package repository

import (
	"blog-v2/src/adapters/repository/inmem"
	"blog-v2/src/testhelpers/random"
	"context"
	"io/fs"
	"testing"
	"testing/fstest"

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
	t.Run("get data from a file", func(t *testing.T) {
		ctx := s.MakeContext(t)
		repo := s.NewRepo

		dir := s.MakeDir(t)

		entries, err := fs.ReadDir(dir, ".")
		assert.NoError(t, err)

		for _, entry := range entries {
			got, found, err := repo.Get(ctx, entry.Name())
			assert.NoError(t, err)
			assert.True(t, found)

			expected, err := fs.ReadFile(dir, entry.Name())
			assert.NoError(t, err)
			assert.Equal(t, expected, got)
		}
	})

	t.Run("return an error if there are no entries in the file system", func(t *testing.T) {
		ctx := s.MakeContext(t)
		repo := s.NewRepo

		dir := fstest.MapFS{}

		entries, err := fs.ReadDir(dir, ".")
		assert.NoError(t, err)

		for _, entry := range entries {
			_, found, err := repo.Get(ctx, entry.Name())
			assert.Error(t, err)
			assert.Equal(t, inmem.ErrEntryNotFound, err)
			assert.False(t, found)
		}
	})

	t.Run("return an error if no relevant entries in the file system", func(t *testing.T) {
		ctx := s.MakeContext(t)
		repo := s.NewRepo

		dir := random.DirFS("blah")

		entries, err := fs.ReadDir(dir, ".")
		assert.NoError(t, err)

		for _, entry := range entries {
			_, found, err := repo.Get(ctx, entry.Name())
			assert.Error(t, err)
			assert.Equal(t, inmem.ErrEntryNotFound, err)
			assert.False(t, found)
		}
	})
}
