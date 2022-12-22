package blog

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
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

	post, err := convertToPost(bytes.NewReader(stuff))
	if err != nil {
		return Post{}, fmt.Errorf("error converting to a blog post: %v", err)
	}

	return post, nil
}

func convertToPost(r io.Reader) (Post, error) {
	post := Post{}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title := readLine()
	post.Title = title
	date, err := StringToDate(readLine())
	if err != nil {
		return Post{}, err
	}
	post.Date = date
	post.Picture = readLine()
	post.Tags = strings.Split(readLine(), ",")
	readLine()

	body := bytes.Buffer{}
	for scanner.Scan() {
		body.Write(scanner.Bytes())
		body.WriteString("\n")
	}
	post.Content = RenderMarkdown(body.Bytes())
	post.URLTitle = strings.Replace(title, " ", "-", -1)

	return post, nil
}
