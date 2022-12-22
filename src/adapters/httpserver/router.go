package httpserver

import (
	"blog-v2/src/adapters/httpserver/bloghandler"
	"blog-v2/src/adapters/httpserver/healthcheckhandler"
	"embed"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed bloghandler/css/*
var css embed.FS

func NewRouter(
	blogService bloghandler.BlogService,
) *mux.Router {
	blogHandler := bloghandler.NewHandler(blogService)

	router := mux.NewRouter()
	router.HandleFunc("/internal/healthcheck", healthcheckhandler.HealthCheck)
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", blogHandler.Public(css)))

	router.Handle("/blog/{title}", http.HandlerFunc(blogHandler.Read))
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
