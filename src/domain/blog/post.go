package blog

import (
	"html/template"
	"time"
)

type Post struct {
	Title    string
	Content  template.HTML
	Date     time.Time
	Picture  string
	Tags     []string
	URLTitle string
}
