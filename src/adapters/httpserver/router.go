package httpserver

import (
	"blog-v2/src/adapters/httpserver/healthcheckhandler"
	"net/http"

	bloghandler "blog-v2/src/adapters/httpserver/bloghandler"

	"github.com/gorilla/mux"
)

const (
	blogPath        = "/blog/{title}"
	healthCheckPath = "/internal/healthcheck"
)

func NewRouter(
	blogService bloghandler.BlogService,
) *mux.Router {
	blogHandler := bloghandler.NewHandler(blogService)

	router := mux.NewRouter()
	router.HandleFunc(healthCheckPath, healthcheckhandler.HealthCheck)
	router.Handle(blogPath, http.HandlerFunc(blogHandler.Read))
	return router
}
