package greethandler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//go:generate moq-v0.2.7 -stub -out greeterservice_moq.go . GreeterService
type GreeterService interface {
	Greet(ctx context.Context, name string) (greeting string, err error)
}

type GreeterServiceFunc func(context.Context, string) (string, error)

func (g GreeterServiceFunc) Greet(ctx context.Context, name string) (greeting string, err error) {
	return g(ctx, name)
}

type GreetHandler struct {
	greeter GreeterService
}

func NewGreetHandler(
	greeter GreeterService,
) *GreetHandler {
	return &GreetHandler{
		greeter: greeter,
	}
}

func (g *GreetHandler) Greet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := mux.Vars(r)["name"]

	if strings.Contains(name, "error") {
		// logger.Error(ctx, "Name contains error!")
		return
	}

	greeting, err := g.greeter.Greet(ctx, name)
	if err != nil {
		// logger.Error(ctx, "failed to greet", zap.Error(err), zap.String("name", name))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprint(w, greeting)
}
