package random

import (
	"blog-v2/src/domain/blog"
	"blog-v2/src/testhelpers/random"
)

func Post() blog.Post {
	return blog.Post{
		Title:   random.StringWithPrefix("title"),
		Content: random.StringWithPrefix("content"),
	}
}
