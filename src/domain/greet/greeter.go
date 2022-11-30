package greet

import (
	"context"
	"fmt"
)

func HelloGreeter(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
