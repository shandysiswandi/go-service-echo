package posts

import (
	"bytes"
	"encoding/json"
	"go-rest-echo/external/jsonplaceholder"
	"net/http"
)

// Create is
func Create(data Post) (*Post, error) {
	post := new(Post)

	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, jsonplaceholder.ErrPostCreate
	}

	response, err := http.Post(jsonplaceholder.PostsPath, "application/json; charset=utf-8", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, jsonplaceholder.ErrPostCreate
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(post); err != nil {
		return nil, jsonplaceholder.ErrPostCreate
	}

	return post, nil
}
