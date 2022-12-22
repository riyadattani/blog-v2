package random

import "testing/fstest"

func DirFS(titles ...string) fstest.MapFS {
	dirFS := make(fstest.MapFS)
	for _, title := range titles {
		dirFS[title] = &fstest.MapFile{Data: []byte(String())}
	}
	return dirFS
}

func DirFSHardcoded() fstest.MapFS {
	return fstest.MapFS{
		"first-post.md":  {Data: []byte(firstPost)},
		"second-post.md": {Data: []byte(secondPost)},
	}
}

const (
	firstPost = `This is the title
2013-Mar-03
picture.jpg
cat,dog
-----
blah blah blah`

	secondPost = `This is the title of the second post
2013-Mar-01
picture2.jpg
bird,fly
-----
This is the first sentence of the content.`
)
