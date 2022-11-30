package httpserver

import (
	"blog-v2/src/adapters/httpserver/greethandler"
	"blog-v2/src/adapters/httpserver/healthcheckhandler"
	"net/http"

	bloghandler "blog-v2/src/adapters/httpserver/blogHandler"

	"github.com/gorilla/mux"
)

const (
	greetPath       = "/greet/{name}"
	blogPath        = "/blog/{title}"
	healthCheckPath = "/internal/healthcheck"
)

func NewRouter(
	greeter greethandler.GreeterService,
	blogService bloghandler.BlogService,
) *mux.Router {
	greetingHandler := greethandler.NewGreetHandler(greeter)
	blogHandler := bloghandler.NewHandler(blogService)

	router := mux.NewRouter()
	router.HandleFunc(healthCheckPath, healthcheckhandler.HealthCheck)
	router.Handle(blogPath, http.HandlerFunc(blogHandler.Read))
	router.Handle(greetPath, http.HandlerFunc(greetingHandler.Greet))
	return router
}
