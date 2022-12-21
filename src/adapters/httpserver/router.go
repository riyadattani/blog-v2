package httpserver

import (
	"blog-v2/src/adapters/httpserver/bloghandler"
	"blog-v2/src/adapters/httpserver/healthcheckhandler"
	"embed"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	blogByTitlePath = "/blog/{title}"
	blogPath        = "/blog"
	healthCheckPath = "/internal/healthcheck"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func NewRouter(
	blogService bloghandler.BlogService,
) *mux.Router {
	blogHandler := bloghandler.NewHandler(blogService)

	router := mux.NewRouter()
	router.HandleFunc(healthCheckPath, healthcheckhandler.HealthCheck)

	router.Handle(blogByTitlePath, http.HandlerFunc(blogHandler.Read))
	router.Handle(blogPath, http.HandlerFunc(blogHandler.Publish))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	return router
}
