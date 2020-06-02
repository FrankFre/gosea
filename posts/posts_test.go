package posts

import (
	"testing"
)

func TestPosts_loadPosts(t *testing.T) {
	p := NewWithSea()

	posts, err := p.LoadPosts()

	t.Log(err)
	t.Log(posts)
}