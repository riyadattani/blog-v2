//go:build unit

package blog_test

import (
	"blog-v2/src/domain/blog"
	"blog-v2/src/specifications"
	"context"
	"testing"
)

func TestBlog(t *testing.T) {
	specifications.Blog{
		Subject: blog.New(),
		MakeCTX: func(tb testing.TB) context.Context {
			return context.Background()
		},
	}.Test(t)
}
