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
		"first-post.md":  {Data: []byte("blah")},
		"second-post.md": {Data: []byte("blah blah")},
	}
}
