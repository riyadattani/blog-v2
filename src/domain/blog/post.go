package blog

import (
	"html/template"
	"time"
)

type Post struct {
	Title    string `json:"title"`
	Content  template.HTML
	Date     time.Time
	Picture  string
	Tags     []string `json:"tags"`
	URLTitle string   `json:"url_title"`
}
