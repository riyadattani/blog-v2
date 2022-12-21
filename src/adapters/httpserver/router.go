package httpserver

import (
	"blog-v2/src/adapters/httpserver/bloghandler"
	"blog-v2/src/adapters/httpserver/healthcheckhandler"
	"embed"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	eventsPath = "/events"
)

//go:embed bloghandler/css/*
var css embed.FS

func NewRouter(
	blogService bloghandler.BlogService,
) *mux.Router {
	blogHandler := bloghandler.NewHandler(blogService)

	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", blogHandler.Public(css)))
	router.HandleFunc("/internal/healthcheck", healthcheckhandler.HealthCheck)
	router.Handle("/blog/{title}", http.HandlerFunc(blogHandler.Read))
	router.Handle("/blog", http.HandlerFunc(blogHandler.Publish)).Methods(http.MethodPost)
	router.Handle("/about", http.HandlerFunc(blogHandler.About))
	router.Handle("/events", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, "Under construction :)")
	}))

	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, "Under construction :)")
	}))

	return router
}
