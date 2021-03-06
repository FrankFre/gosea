package posts

import "encoding/json"

type RemotePost struct {
	UserID json.Number `json:"userId"`
	ID     json.Number `json:"id"`
	Title  string      `json:"title,omitempty"`
	Body   string      `json:"body,omitempty"`
}

type ResponsePost struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Title    string `json:"title,omitempty"`
	Body     string `json:"body,omitempty"`
}

