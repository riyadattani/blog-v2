package acceptance

import (
	"blog-v2/src"
	"blog-v2/src/adapters/httpserver"
	"blog-v2/src/domain/blog"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

type APIClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewAPIClient(transport http.RoundTripper, baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout:   5 * time.Second,
			Transport: transport,
		},
	}
}

func (a *APIClient) ReadPost(ctx context.Context, title string) (blog.Post, error) {
	fs := os.DirFS("testdata")
	newApp := src.NewApp(ctx, fs)

	server := httpserver.NewWebServer(httpserver.ServerConfig{}, httpserver.NewRouter(newApp.BlogService))
	svr := httptest.NewServer(server.Handler)

	defer svr.Close()

	url := svr.URL + "/blog/" + title

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

func (a *APIClient) SavePost(ctx context.Context, post os.File) error {
	url := a.baseURL + "/blog"

	fileBytes, err := os.Open(post.Name())
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, fileBytes)
	if err != nil {
		return fmt.Errorf("could not create new request: %w", err)
	}

	res, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("problem reaching %s: %w", url, err)
	}

	defer res.Body.Close()
	all, _ := io.ReadAll(res.Body)

	if res.StatusCode == http.StatusInternalServerError {
		return fmt.Errorf(string(all))
	}

	if res.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unexpected status %d from POST %q, body: %q", res.StatusCode, url, string(all))
	}

	return nil
}

func (a *APIClient) WaitForAPIToBeHealthy(ctx context.Context, retries int) error {
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

func (a *APIClient) checkIfHealthy(ctx context.Context) error {
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
