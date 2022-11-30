package httpserver_test

import (
	"blog-v2/black-box-tests/acceptance"
	"blog-v2/src/adapters/httpserver"
	"blog-v2/src/adapters/httpserver/greethandler"
	"blog-v2/src/domain/greet"
	"blog-v2/src/specifications"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewWebServer(t *testing.T) {
	router := httpserver.NewRouter(
		greethandler.GreeterServiceFunc(greet.HelloGreeter),
	)
	webServer := httpserver.NewWebServer(
		httpserver.ServerConfig{},
		router,
	)

	svr := httptest.NewServer(webServer.Handler)
	defer svr.Close()

	client := acceptance.NewAPIClient(http.DefaultTransport, svr.URL)

	specifications.Greeting{
		Subject: client,
		MakeContext: func(tb testing.TB) context.Context {
			return context.Background()
		},
	}.Test(t)
}
