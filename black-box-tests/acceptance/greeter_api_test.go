//go:build acceptance

package acceptance

import (
	"blog-v2/src/specifications"
	"context"
	"net/http"
	"testing"

	"github.com/alecthomas/assert/v2"
)

const fiveRetries = 5

func TestGreetingApplication(t *testing.T) {
	config, err := LoadTestingConfig()
	assert.NoError(t, err)

	client := NewAPIClientGreeter(http.DefaultTransport, config.BaseURL)

	if err := client.WaitForAPIToBeHealthy(context.Background(), fiveRetries); err != nil {
		t.Fatal(err)
	}

	t.Run("api can do greetings", func(t *testing.T) {
		specifications.Greeting{
			Subject: client,
			MakeContext: func(tb testing.TB) context.Context {
				return context.Background()
			},
		}.Test(t)
	})
}
