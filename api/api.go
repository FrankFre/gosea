package api

import (
	"encoding/json"
	"github.com/FrankFre/gosea/posts"
	"net/http"
)

type postsService interface {
	LoadPosts() ([]posts.RemotePost, error)
}

type Api struct {
	posts postsService
}

func New(posts *posts.Posts) *Api {
	return &Api{
		posts: posts,
	}
}

// Posts returns a json response with remote posts
func (a *Api) Posts(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	remotePosts, err := a.posts.LoadPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(remotePosts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
