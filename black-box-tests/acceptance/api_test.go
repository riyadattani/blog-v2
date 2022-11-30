//go:build acceptance

package acceptance_test

import (
	"blog-v2/black-box-tests/acceptance"
	"blog-v2/src/specifications"
	"context"
	"net/http"
	"testing"

	"github.com/alecthomas/assert/v2"
)

const fiveRetries = 5

func TestBlogApplication(t *testing.T) {
	config, err := acceptance.LoadTestingConfig()
	assert.NoError(t, err)

	client := acceptance.NewAPIClient(http.DefaultTransport, config.BaseURL)

	if err := client.WaitForAPIToBeHealthy(context.Background(), fiveRetries); err != nil {
		t.Fatal(err)
	}

	t.Run("api can read a blog", func(t *testing.T) {
		specifications.Blog{
			Subject: client,
			MakeCTX: func(tb testing.TB) context.Context {
				return context.Background()
			},
		}.Test(t)
	})
}
