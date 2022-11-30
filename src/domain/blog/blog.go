package blog

import (
	"context"
)

type Blog struct{}

func New() *Blog {
	return &Blog{}
}

func (b Blog) Publish(ctx context.Context, post Post) error {
	// TODO implement me
	panic("implement me")
}

func (b Blog) Read(ctx context.Context, title string) (post Post, err error) {
	return Post{
		Title:   "life",
		Content: "yo",
	}, nil
}
