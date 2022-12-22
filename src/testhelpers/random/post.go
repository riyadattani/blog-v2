package random

import (
	"blog-v2/src/domain/blog"
	"time"
)

func Post() blog.Post {
	return blog.Post{
		Title:    "This is the title of the builder blog",
		Content:  "This is the content of the builder blog",
		Date:     time.Now(),
		Picture:  "picture.png",
		Tags:     []string{"bob", "builder"},
		URLTitle: "This-is-the-title-of-the-builder-blog",
	}
}
