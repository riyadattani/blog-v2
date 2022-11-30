package specifications

import (
	"blog-v2/src/domain/blog"
	"blog-v2/src/domain/random"
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

	t.Run("can publish and read a blog post", func(t *testing.T) {
		ctx := b.MakeCTX(t)

		postToPublish := random.Post()

		assert.NoError(t, b.Subject.Publish(ctx, postToPublish))

		gotPost, err := b.Subject.Read(ctx, postToPublish.Title)
		assert.NoError(t, err)
		assert.Equal(t, postToPublish.Title, gotPost.Title)
		assert.Equal(t, postToPublish.Content, gotPost.Content)
	})
}
