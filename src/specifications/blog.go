package specifications

import (
	"blog-v2/src/domain/blog"
	"blog-v2/src/testhelpers/random"
	"context"
	"io/fs"
	"testing"

	"github.com/alecthomas/assert/v2"
)

type BlogDriver interface {
	ReadPost(ctx context.Context, title string) (blog.Post, error)
}

type Blog struct {
	Subject     BlogDriver
	MakeCTX     func(tb testing.TB) context.Context
	MakeBlogDir func(tb testing.TB) fs.FS
}

func (b Blog) Test(t *testing.T) {
	t.Helper()

	t.Run("can create a post, save it and read it", func(t *testing.T) {
		ctx := b.MakeCTX(t)
		dir := b.MakeBlogDir(t)

		// pp.PP(dir)

		entries, err := fs.ReadDir(dir, ".")
		assert.NoError(t, err)

		for _, entry := range entries {
			post, err := b.Subject.ReadPost(ctx, entry.Name())
			assert.NoError(t, err)

			bytes, err := fs.ReadFile(dir, entry.Name())
			assert.NoError(t, err)

			markdown := string(bytes)
			assert.Contains(t, markdown, post.Title)
			assert.Contains(t, markdown, post.Picture)

			tags := post.Tags
			for _, tag := range tags {
				assert.Contains(t, markdown, tag)
			}

			// assert.Contains(t, markdown, string(post.Content))
			// assert.Contains(t, markdown, post.URLTitle)
			// assert.Contains(t, markdown, post.Date)
		}
	})

	t.Run("error if blog post does not exist", func(t *testing.T) {
		ctx := b.MakeCTX(t)
		dir := random.DirFS("does-not-exist")

		// pp.PP(dir)

		entries, err := fs.ReadDir(dir, ".")
		assert.NoError(t, err)

		for _, entry := range entries {
			_, err := b.Subject.ReadPost(ctx, entry.Name())
			assert.Error(t, err)
		}
	})
}
