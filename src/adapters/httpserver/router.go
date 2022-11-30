package httpserver

import (
	"blog-v2/src/adapters/httpserver/greethandler"
	"blog-v2/src/adapters/httpserver/healthcheckhandler"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	greetPath       = "/greet/{name}"
	healthCheckPath = "/internal/healthcheck"
)

func NewRouter(
	greeter greethandler.GreeterService,
) *mux.Router {
	greetingHandler := greethandler.NewGreetHandler(greeter)

	router := mux.NewRouter()
	router.HandleFunc(healthCheckPath, healthcheckhandler.HealthCheck)
	router.Handle(greetPath, http.HandlerFunc(greetingHandler.Greet))
	return router
}
