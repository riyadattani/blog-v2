package bloghandler

import (
	"blog-v2/src/domain/blog"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go:generate moq-v0.2.7 -stub -out blogservice_moq.go . BlogService
type BlogService interface {
	Publish(ctx context.Context, post blog.Post) error
	Read(ctx context.Context, title string) (post blog.Post, err error)
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
		http.Error(w, "empty title", http.StatusInternalServerError)
		return
	}

	post, err := g.blogService.Read(ctx, title)
	if err != nil {
		log.Println("failed to read post", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(post)
}

func (g *BlogHandler) Publish(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b blog.Post

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "could decode post", http.StatusInternalServerError)
		return
	}

	if err := g.blogService.Publish(ctx, b); err != nil {
		http.Error(w, "could not publish post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
