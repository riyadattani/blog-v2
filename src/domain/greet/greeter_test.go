package greet

import (
	"blog-v2/src/specifications"
	"context"
	"testing"
)

func TestHelloGreeter(t *testing.T) {
	specifications.Greeting{
		Subject: specifications.GreetingSystemFunc(HelloGreeter),
		MakeContext: func(tb testing.TB) context.Context {
			return context.Background()
		},
	}.Test(t)
}
