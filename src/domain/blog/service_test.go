//go:build unit

package blog_test

import (
	"blog-v2/src/adapters/repository/inmem"
	"blog-v2/src/domain/blog"
	"blog-v2/src/specifications"
	"blog-v2/src/testhelpers/random"
	"context"
	"io/fs"
	"testing"
)

func TestBlog(t *testing.T) {
	dirFS := random.DirFSHardcoded()

	repo := inmem.NewRepository(dirFS)
	service := blog.NewService(repo)

	specifications.Blog{
		Subject: service,
		MakeCTX: func(tb testing.TB) context.Context {
			return context.Background()
		},
		MakeBlogDir: func(tb testing.TB) fs.FS {
			return dirFS
		},
	}.Test(t)
}
