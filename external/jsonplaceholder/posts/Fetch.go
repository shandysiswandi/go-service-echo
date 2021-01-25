package posts

import (
	"encoding/json"
	"go-rest-echo/external/jsonplaceholder"
	"net/http"
)

// Fetch is
func Fetch() (*[]Post, error) {
	post := new([]Post)

	response, err := http.Get(jsonplaceholder.PostsPath)
	if err != nil {
		return nil, jsonplaceholder.ErrPostFetch
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(post); err != nil {
		return nil, jsonplaceholder.ErrPostFetch
	}

	return post, nil
}
