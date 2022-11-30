//go:build unit

package blog_test

import (
	"blog-v2/src/adapters/repository/inmem"
	"blog-v2/src/domain/blog"
	"blog-v2/src/specifications"
	"context"
	"testing"
)

func TestBlog(t *testing.T) {
	repo := inmem.NewRepository()
	specifications.Blog{
		Subject: blog.NewService(repo),
		MakeCTX: func(tb testing.TB) context.Context {
			return context.Background()
		},
	}.Test(t)
}
