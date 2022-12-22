package blog

import (
	"context"
	"fmt"
)

type Repository interface {
	Get(ctx context.Context, title string) (stuff []byte, found bool, err error)
}

type Service struct {
	Repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		Repository: repo,
	}
}

func (b Service) ReadPost(ctx context.Context, title string) (Post, error) {
	stuff, found, err := b.Repository.Get(ctx, title)
	if err != nil {
		return Post{}, err
	}

	if !found {
		return Post{}, fmt.Errorf("could not find blog with title %q", title)
	}

	post := convertToPost(stuff)
	if err != nil {
		return Post{}, fmt.Errorf("error converting to a blog post: %v", err)
	}

	return post, nil
}

func convertToPost(_ []byte) Post {
	return Post{
		Title:   "first-post.md",
		Content: "blah",
	}
}
