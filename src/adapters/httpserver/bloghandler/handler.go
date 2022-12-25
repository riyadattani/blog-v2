package bloghandler

import (
	"blog-v2/src/domain/blog"
	"context"
	"embed"
	"encoding/json"
	"errors"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go:generate moq-v0.2.7 -stub -out blogservice_moq.go . BlogService
type BlogService interface {
	ReadPost(ctx context.Context, title string) (post blog.Post, err error)
}

type BlogHandler struct {
	blogService BlogService
}

func NewHandler(
	blogService BlogService,
) *BlogHandler {
	return &BlogHandler{
		blogService: blogService,
	}
}

func (g *BlogHandler) Read(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	title := mux.Vars(r)["title"]
	if title == "" {
		http.Error(w, "empty title", http.StatusBadRequest)
		return
	}

	post, err := g.blogService.ReadPost(ctx, title)
	var errNotFound blog.ErrNotFound
	if errors.As(err, &errNotFound) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Println("failed to read post: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "error encoding post to JSON", http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "blog.gohtml", post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//go:embed templates/*
var templates embed.FS

var t = template.Must(template.ParseFS(templates, "templates/*"))

func (g *BlogHandler) About(w http.ResponseWriter, r *http.Request) {
	if err := t.ExecuteTemplate(w, "about.gohtml", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (g *BlogHandler) Public(css fs.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := fs.Sub(css, "bloghandler/css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.FileServer(http.FS(f)).ServeHTTP(w, r)
	})
}
