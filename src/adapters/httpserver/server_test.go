package httpserver_test

import (
	"blog-v2/src/adapters/httpserver"
	"blog-v2/src/adapters/repository/inmem"
	"blog-v2/src/domain/blog"
	"blog-v2/src/specifications"
	"context"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"
)

func TestNewWebServer(t *testing.T) {
	dirFS := fstest.MapFS{
		"first-post.md":  {Data: []byte("blah")},
		"second-post.md": {Data: []byte("blah blah")},
	}

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
