//go:build unit

package inmem_test

import (
	"blog-v2/src/adapters/repository"
	"blog-v2/src/adapters/repository/inmem"
	"context"
	"testing"
)

func TestRepository(t *testing.T) {
	repository.Specification{
		NewRepo: inmem.NewRepository(),
		MakeContext: func(tb testing.TB) context.Context {
			return context.Background()
		},
	}.Test(t)
}
