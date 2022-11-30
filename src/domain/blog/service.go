package blog

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, title string) (blog Post, found bool, err error)
	Create(ctx context.Context, blog Post) error
}

type Service struct {
	Repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		Repository: repo,
	}
}

func (b Service) Publish(ctx context.Context, post Post) error {
	// TODO implement me
	panic("implement me")
}

func (b Service) Read(ctx context.Context, title string) (post Post, err error) {
	return Post{
		Title:   "life",
		Content: "yo",
	}, nil
}
