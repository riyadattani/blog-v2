package inmem

import (
	"blog-v2/src/domain/blog"
	"context"
)

type Repository struct {
	posts []blog.Post
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Get(ctx context.Context, title string) (blog blog.Post, found bool, err error) {
	for _, b := range r.posts {
		if b.Title == title {
			return b, true, nil
		}
	}

	return blog, false, nil
}

func (r *Repository) Create(ctx context.Context, post blog.Post) error {
	r.posts = append(r.posts, post)
	return nil
}
