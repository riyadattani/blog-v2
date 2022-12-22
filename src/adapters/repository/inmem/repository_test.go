//go:build unit

package inmem_test

import (
	"blog-v2/src/adapters/repository"
	"blog-v2/src/adapters/repository/inmem"
	"context"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestRepository(t *testing.T) {
	dirFS := fstest.MapFS{
		"first-post.md":  {Data: []byte("blah")},
		"second-post.md": {Data: []byte("blah blah")},
	}

	repository.Specification{
		NewRepo: inmem.NewRepository(dirFS),
		MakeContext: func(tb testing.TB) context.Context {
			return context.Background()
		},

		MakeDir: func(tb testing.TB) fs.FS {
			return dirFS
		},
	}.Test(t)
}
