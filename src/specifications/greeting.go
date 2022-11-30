package specifications

import (
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
)

type GreetingSystemDriver interface {
	Greet(ctx context.Context, name string) (greeting string, err error)
}

type Greeting struct {
	Subject     GreetingSystemDriver
	MakeContext func(tb testing.TB) context.Context
}

func (s Greeting) Test(t *testing.T) {
	t.Helper()

	t.Run("greets people in a friendly manner", func(t *testing.T) {
		ctx := s.MakeContext(t)
		greeting, err := s.Subject.Greet(ctx, "Pepper")
		assert.NoError(t, err)
		assert.Equal(t, "Hello, Pepper!", greeting)
	})
}

type GreetingSystemFunc func(ctx context.Context, name string) (greeting string, err error)

func (g GreetingSystemFunc) Greet(ctx context.Context, name string) (greeting string, err error) {
	return g(ctx, name)
}
