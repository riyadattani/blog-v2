package httpserver_test

import (
	"blog-v2/src/adapters/httpserver"
	"blog-v2/src/adapters/repository/inmem"
	"blog-v2/src/domain/blog"
	"blog-v2/src/specifications"
	"blog-v2/src/testhelpers/random"
	"context"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewWebServer(t *testing.T) {
	dirFS := random.DirFSHardcoded()
	router := httpserver.NewRouter(
		blog.NewService(inmem.NewRepository(dirFS)),
	)
	webServer := httpserver.NewWebServer(
		httpserver.ServerConfig{},
		router,
	)

	svr := httptest.NewServer(webServer.Handler)
	defer svr.Close()

	specifications.Blog{
		Subject: httpserver.NewClient(http.DefaultTransport, svr.URL),
		MakeCTX: func(tb testing.TB) context.Context {
			return context.Background()
		},
		MakeBlogDir: func(tb testing.TB) fs.FS {
			return dirFS
		},
	}.Test(t)
}
