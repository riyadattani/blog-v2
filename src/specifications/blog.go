package specifications

import (
	"blog-v2/src/domain/blog"
	"context"
	"io/fs"
	"testing"

	"github.com/adamluzsi/testcase/pp"

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

		blogDir := b.MakeBlogDir(t)

		pp.PP(blogDir)

		entries, err := fs.ReadDir(blogDir, ".")
		assert.NoError(t, err)

		for _, entry := range entries {
			_, err := b.Subject.ReadPost(ctx, entry.Name())
			assert.NoError(t, err)

			t.Logf(entry.Name())

			// TODO: read the content!

			//f, err := blogDir.Open(entry.Name())
			//assert.NoError(t, err)
			//
			//defer f.Close()
		}
	})
}
