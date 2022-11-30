package blog

import (
	"context"
	"fmt"
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
	if err := b.Repository.Create(ctx, post); err != nil {
		return err
	}
	return nil
}

func (b Service) Read(ctx context.Context, title string) (post Post, err error) {
	post, found, err := b.Repository.Get(ctx, title)
	if err != nil {
		return Post{}, err
	}

	if !found {
		return Post{}, fmt.Errorf("could not find blog with title %q", title)
	}
	return post, nil
}
