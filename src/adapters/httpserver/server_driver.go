package httpserver

import (
	"blog-v2/src/domain/blog"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(transport http.RoundTripper, baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout:   5 * time.Second,
			Transport: transport,
		},
	}
}

func (a *Client) ReadPost(ctx context.Context, title string) (blog.Post, error) {
	url := a.baseURL + "/blog/" + title

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return blog.Post{}, fmt.Errorf("could not create request - %w", err)
	}

	res, err := a.httpClient.Do(req)
	if err != nil {
		return blog.Post{}, fmt.Errorf("problem reaching %s, %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return blog.Post{}, fmt.Errorf("unexpected status %d from GET %q", res.StatusCode, url)
	}

	var response blog.Post

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return blog.Post{}, fmt.Errorf("problem decoding JSON: %w", err)
	}

	return response, nil
}

func (a *Client) WaitForAPIToBeHealthy(ctx context.Context, retries int) error {
	var (
		err   error
		start = time.Now()
	)

	for retries > 0 {
		if err = a.checkIfHealthy(ctx); err != nil {
			retries -= 1
			time.Sleep(1 * time.Second)
		} else {
			return nil
		}
	}

	return fmt.Errorf("given up checking healthcheck after %dms, %v", time.Since(start).Milliseconds(), err)
}

func (a *Client) checkIfHealthy(ctx context.Context) error {
	url := a.baseURL + "/internal/healthcheck"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("could not create request - %w", err)
	}

	res, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("problem reaching %s, %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d from POST %q", res.StatusCode, url)
	}

	return nil
}
