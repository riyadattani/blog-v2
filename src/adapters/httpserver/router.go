package httpserver

import (
	"blog-v2/src/adapters/httpserver/bloghandler"
	"blog-v2/src/adapters/httpserver/healthcheckhandler"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	blogByTitlePath = "/blog/{title}"
	blogPath        = "/blog"
	healthCheckPath = "/internal/healthcheck"
)

func NewRouter(
	blogService bloghandler.BlogService,
) *mux.Router {
	blogHandler := bloghandler.NewHandler(blogService)

	router := mux.NewRouter()
	router.HandleFunc(healthCheckPath, healthcheckhandler.HealthCheck)

	router.Handle(blogByTitlePath, http.HandlerFunc(blogHandler.Read))
	router.Handle(blogPath, http.HandlerFunc(blogHandler.Publish))
	return router
}
