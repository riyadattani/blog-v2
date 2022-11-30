package httpserver_test

import (
	"blog-v2/black-box-tests/acceptance"
	"blog-v2/src/adapters/httpserver"
	"blog-v2/src/adapters/repository/inmem"
	"blog-v2/src/domain/blog"
	"blog-v2/src/specifications"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewWebServer(t *testing.T) {
	router := httpserver.NewRouter(
		blog.NewService(inmem.NewRepository()),
	)
	webServer := httpserver.NewWebServer(
		httpserver.ServerConfig{},
		router,
	)

	svr := httptest.NewServer(webServer.Handler)
	defer svr.Close()

	specifications.Blog{
		Subject: acceptance.NewAPIClient(http.DefaultTransport, svr.URL),
		MakeCTX: func(tb testing.TB) context.Context {
			return context.Background()
		},
	}.Test(t)
}
