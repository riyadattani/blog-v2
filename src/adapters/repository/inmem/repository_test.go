//go:build unit

package inmem_test

import (
	"blog-v2/src/adapters/repository"
	"blog-v2/src/adapters/repository/inmem"
	"blog-v2/src/testhelpers/random"
	"context"
	"io/fs"
	"testing"
)

func TestRepository(t *testing.T) {
	dirFS := random.DirFSHardcoded()
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
