package specifications

import (
	"blog-v2/src/domain/blog"
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
)

type BlogDriver interface {
	Publish(ctx context.Context, post blog.Post) error
	Read(ctx context.Context, title string) (post blog.Post, err error)
}

type Blog struct {
	Subject BlogDriver
	MakeCTX func(tb testing.TB) context.Context
}

func (b Blog) Test(t *testing.T) {
	t.Helper()

	t.Run("can read a blog post", func(t *testing.T) {
		ctx := b.MakeCTX(t)
		post, err := b.Subject.Read(ctx, "life")

		assert.NoError(t, err)
		assert.Equal(t, "life", post.Title)
		assert.Equal(t, "yo", post.Content)
	})
}
