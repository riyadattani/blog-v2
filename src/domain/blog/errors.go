package blog

import "fmt"

type ErrNotFound struct {
	Title string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("could not find blog post with title: %v", e.Title)
}
