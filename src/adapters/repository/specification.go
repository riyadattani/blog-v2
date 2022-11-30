package repository

import (
	"blog-v2/src/domain/blog"
	"blog-v2/src/domain/random"
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Driver interface {
	Get(ctx context.Context, title string) (blog blog.Post, found bool, err error)
	Create(ctx context.Context, blog blog.Post) error
	// Update(ctx context.Context, blog blog.Post) error
	// Delete(ctx context.Context, title string) error
}

type Specification struct {
	NewRepo     Driver
	MakeContext func(tb testing.TB) context.Context
}

func (s Specification) Test(t *testing.T) {
	t.Run("Create post", func(t *testing.T) {
		t.Run("it cannot create a post with the same title twice", func(t *testing.T) {
		})
	})
	t.Run("Get post by title", func(t *testing.T) {
		t.Run("it can retrieve a stored post by title", func(t *testing.T) {
			ctx := s.MakeContext(t)
			post := random.Post()

			repo := s.NewRepo
			assert.NoError(t, repo.Create(ctx, post))
			got, found, err := repo.Get(ctx, post.Title)

			assert.NoError(t, err)
			assert.True(t, found)
			assert.Equal(t, post, got)
		})

		t.Run("it returns an empty slice when there are not posts by that title", func(t *testing.T) {
		})
	})
}
