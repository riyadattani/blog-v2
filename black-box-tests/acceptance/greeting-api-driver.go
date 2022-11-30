package acceptance

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIClientGreeter struct {
	baseURL    string
	httpClient *http.Client
}

func NewAPIClientGreeter(transport http.RoundTripper, baseURL string) *APIClientGreeter {
	return &APIClientGreeter{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout:   5 * time.Second,
			Transport: transport,
		},
	}
}

func (a *APIClientGreeter) Greet(ctx context.Context, name string) (string, error) {
	url := a.baseURL + "/greet/" + name

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("could not create request - %w", err)
	}

	res, err := a.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("problem reaching %s, %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d from GET %q", res.StatusCode, url)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (a *APIClientGreeter) checkIfHealthy(ctx context.Context) error {
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

func (a *APIClientGreeter) WaitForAPIToBeHealthy(ctx context.Context, retries int) error {
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
