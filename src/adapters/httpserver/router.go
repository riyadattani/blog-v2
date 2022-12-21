package httpserver

import (
	"blog-v2/src/adapters/httpserver/bloghandler"
	"blog-v2/src/adapters/httpserver/healthcheckhandler"
	"embed"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	blogByTitlePath = "/blog/{title}"
	blogPath        = "/blog"
	healthCheckPath = "/internal/healthcheck"
)

//go:embed bloghandler/css/*
var css embed.FS

func NewRouter(
	blogService bloghandler.BlogService,
) *mux.Router {
	blogHandler := bloghandler.NewHandler(blogService)

	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", blogHandler.Public(css)))
	router.HandleFunc(healthCheckPath, healthcheckhandler.HealthCheck)
	router.Handle(blogByTitlePath, http.HandlerFunc(blogHandler.Read))
	router.Handle(blogPath, http.HandlerFunc(blogHandler.Publish))
	router.Handle("/about", http.HandlerFunc(blogHandler.About))

	return router
}
